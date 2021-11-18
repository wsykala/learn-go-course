/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/
package main

import "fmt"

func main() {
	favSport := "Something"
	switch favSport {
	case "Something":
		fmt.Println("Soo you like Something?")
	case "Football":
		fmt.Println("This is nice too")
	}
}
