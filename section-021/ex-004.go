/*
Fix the race condition you created in the previous exercise by using a mutex
	it makes sense to remove runtime.Gosched()
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	const count = 10
	var wg sync.WaitGroup
	wg.Add(count)

	var mu sync.Mutex

	for i := 0; i < count; i++ {
		go func() {
			mu.Lock()
			x := counter
			runtime.Gosched()
			x++
			counter = x
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
