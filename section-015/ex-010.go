/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:
*/
package main

import "fmt"

func foo() func() {
	var x int
	return func() {
		fmt.Printf("x=%d\n", x)
		x++
	}
}

func main() {
	f := foo()
	g := foo()
	f()
	f()
	f()
	g()
	g()
	f()
}
