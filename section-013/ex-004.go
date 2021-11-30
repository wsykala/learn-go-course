/*
Create and use an anonymous struct.
*/
package main

import "fmt"

func main() {
	p := struct {
		first string
		last  string
	}{
		first: "John",
		last:  "Johnson",
	}

	fmt.Printf("%T %s", p, p)
}
