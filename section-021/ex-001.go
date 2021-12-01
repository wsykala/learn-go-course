/*
in addition to the main goroutine, launch two additional goroutines
	each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Printf("Some goroutine 1\n\tGoroutines %d\n", runtime.NumGoroutine())
	wg.Done()
}

func main() {
	wg.Add(2)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	go foo()
	go func() {
		fmt.Printf("Some goroutine 2\n\tGoroutines %d\n", runtime.NumGoroutine())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
}
