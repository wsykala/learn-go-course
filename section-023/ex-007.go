/*
write a program that
	launches 10 goroutines
		each goroutine adds 10 numbers to a channel
	pull the numbers off the channel and print them
*/
package main

import "fmt"

func main() {
	// Uses the approach with many channels, just for the fun of it.
	// All of this can be done with a single channel
	const count = 10

	receiver := make(chan int)
	channels := make([]<-chan int, 0, count)

	for i := 0; i < count; i++ {
		channels = append(channels, generateData(i))
	}

	go receive(receiver, channels)
	for v := range receiver {
		fmt.Println(v)
	}
}

func generateData(m int) <-chan int {
	ch := make(chan int)
	go func() {
		for j := 0; j < 10; j++ {
			ch <- j * m
		}
		close(ch)
	}()
	return ch
}

func receive(rec chan<- int, ch []<-chan int) {
	for _, c := range ch {
		for v := range c {
			rec <- v
		}
	}
	close(rec)
}
