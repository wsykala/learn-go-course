/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/
package main

import "fmt"

func foo() {
	fmt.Println("Hey, I am inside foo!")
}

func main() {
	defer foo()
	fmt.Println("Foo() will get called after main returns!")
}
