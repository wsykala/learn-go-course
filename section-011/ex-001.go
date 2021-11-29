/*
Using a COMPOSITE LITERAL:
	create an ARRAY which holds 5 VALUES of TYPE int
	assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
	print out the TYPE of the array
*/
package main

import "fmt"

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = 42 + i
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", arr)

}
