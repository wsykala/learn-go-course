/*
Create a program that shows the “if statement” in action.
And ex 007
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/
package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i == 0 {
			fmt.Println("Zero")
		} else if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
