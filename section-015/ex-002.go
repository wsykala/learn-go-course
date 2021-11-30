/*
create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
*/
package main

import "fmt"

func foo(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func bar(numbers []int) int {
	return foo(numbers...)
}

func main() {
	numbers := []int{1, 2, 10, 30, 21, 44, 42}
	fmt.Printf("[foo] The sum of %d is %d\n", numbers, foo(numbers...))
	fmt.Printf("[bar] The sum of %d is %d\n", numbers, bar(numbers))
}
