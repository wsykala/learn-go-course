/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/
package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 44, "Hello"
}

func main() {
	fmt.Println("Calling foo:", foo())
	x, y := bar()
	fmt.Println("Return values from bar:", x, y)
}
