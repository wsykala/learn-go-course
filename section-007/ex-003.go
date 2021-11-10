/*
Create TYPED and UNTYPED constants. Print the values of the constants.
*/

package main

import "fmt"

const (
	a         = 42
	b float64 = 42.42
)

func main() {
	fmt.Println(a, b)
}
