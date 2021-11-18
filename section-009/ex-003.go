/*
Create a for loop using this syntax
	for condition { }
Have it print out the years you have been alive.
*/
package main

import "fmt"

func main() {
	year := 1996
	for year <= 2021 {
		fmt.Println(year)
		year++
	}
}
