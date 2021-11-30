/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
	first name
	last name
	favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/
package main

import "fmt"

type Person struct {
	first    string
	last     string
	flavours []string
}

func main() {
	p1 := Person{
		first: "John",
		last:  "Johnson",
		flavours: []string{
			"Strwaberry",
			"Vanilla",
		},
	}
	p2 := Person{
		first: "Mark",
		last:  "Hogan",
		flavours: []string{
			"Mint",
			"Chocolate",
			"Raspberry",
		},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.flavours {
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.flavours {
		fmt.Printf("%s ", v)
	}
}
