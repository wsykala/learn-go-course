/*
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
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

	pm := map[string]Person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range pm {
		fmt.Println(k, v)
	}
}
