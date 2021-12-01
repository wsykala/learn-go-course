/*
This exercise will reinforce our understanding of method sets:
	create a type person struct
	attach a method speak to type person using a pointer receiver
		*person
	create a type human interface
		to implicitly implement the interface, a human must have the speak method
	create func “saySomething”
		have it take in a human as a parameter
		have it call the speak method
	show the following in your code
		you CAN pass a value of type *person into saySomething
		you CANNOT pass a value of type person into saySomething
*/
package main

import "fmt"

type Person struct {
	First string
	Last  string
}

func (p *Person) speak() {
	fmt.Println("Hello, my name is", p.First, p.Last)
}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func main() {
	p := Person{
		First: "John",
		Last:  "Johnson",
	}

	// The below won't work!
	// saySomething(p)
	saySomething(&p)
}
