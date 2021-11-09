package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	foo()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func foo() {
	fmt.Println("This is how you do a function")
}
