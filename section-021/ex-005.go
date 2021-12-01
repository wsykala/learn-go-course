/*
Fix the race condition you created in exercise #4 by using package atomic
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	const count = 10
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			x := atomic.LoadInt64(&counter)
			x++
			atomic.StoreInt64(&counter, x)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt64(&counter))
}
