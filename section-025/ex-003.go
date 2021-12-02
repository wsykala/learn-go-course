/*
Create a struct “customErr” which implements the builtin.error interface.
Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”.
 If you need a hint, here is one.
*/
package main

import "fmt"

type customErr struct {
}

func (ce customErr) Error() string {
	return "Our custom error"
}

func foo(err error) {
	fmt.Printf("%T\n", err)
}

func main() {
	x := customErr{}
	foo(x)
}
