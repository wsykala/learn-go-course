/*
Create a func which returns a func
assign the returned func to a variable
call the returned func
*/
package main

import "fmt"

func foo() func() bool {
	return func() bool {
		return true
	}
}

func main() {
	f := foo()
	fmt.Println(f())
}
