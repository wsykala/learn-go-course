/*
Assign a func to a variable, then call that func
*/
package main

import "fmt"

func main() {
	f := func(a int) bool {
		return a%2 == 0
	}

	x := 2
	y := 3
	fmt.Printf("[Anonymous func] %d is even %t\n", x, f(x))
	fmt.Printf("[Anonymous func] %d is even %t\n", y, f(y))
}
