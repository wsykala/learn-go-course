/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
	pass a func into a func as an argument
*/
package main

import "fmt"

func even(x ...int) []int {
	var even []int
	for _, v := range x {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return even
}

func sum(f func(x ...int) []int, numbers []int) int {
	sum := 0
	for _, v := range f(numbers...) {
		sum += v
	}
	return sum
}

func main() {
	x := []int{1, 3, 4, 2, 5}
	fmt.Println("Sum of even in", x, "is", sum(even, x))
}
