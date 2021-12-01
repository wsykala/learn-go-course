/*
Using goroutines, create an incrementer program
	have a variable to hold the incrementer value
	launch a bunch of goroutines
		each goroutine should
			read the incrementer value
				store it in a new variable
			yield the processor with runtime.Gosched()
			increment the new variable
			write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition.
Prove that it is a race condition by using the -race flag
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var incrementer int

func main() {
	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			x := incrementer
			runtime.Gosched()
			x++
			incrementer = x
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)
	// Sample result:
	// 10
	// Found 1 data race(s)
	// exit status 66
}
