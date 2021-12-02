/*
get this code working:
	using func literal, aka, anonymous self-executing func
	a buffered channel
*/
package main

import (
	"fmt"
)

func firstSolution() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func secondSolution() {
	c := make(chan int, 1)

	c <- 42
	fmt.Println(<-c)
}

func main() {
	firstSolution()
	secondSolution()
}
