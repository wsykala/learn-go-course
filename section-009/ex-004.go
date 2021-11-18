/*
Create a for loop using this syntax
	for { }
Have it print out the years you have been alive.
*/
package main

import "fmt"

func main() {
	year := 1996
	for {
		if year > 2021 {
			break
		}
		fmt.Println(year)
		year++
	}
}
