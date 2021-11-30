/*
Create a user defined struct with
	the identifier “person”
	the fields:
		first
		last
		age
attach a method to type person with
	the identifier “speak”
	the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/
package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	p := Person{
		first: "John",
		last:  "Johnson",
		age:   42,
	}
	p.speak()
}
