/*
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
	circle area= π r 2
	square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/
package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float32
	width  float32
}

type Circle struct {
	r float32
}

type Shape interface {
	area() float32
}

func (s Square) area() float32 {
	return s.length * s.width
}

func (c Circle) area() float32 {
	return math.Pi * c.r * c.r
}

func info(s Shape) {
	// The switch statement is optional here, but I am leaving it just to remember about the usage
	// The line below will work just fine if we are not interested about custom message
	// fmt.Println(s.area())
	switch s.(type) {
	case Square:
		fmt.Println("The area of a square is", s.(Square).area())
	case Circle:
		fmt.Println("The area of a circle is", s.(Circle).area())
	}
}

func main() {
	sq := Square{
		length: 42,
		width:  42,
	}
	c := Circle{
		r: 20,
	}

	info(sq)
	info(c)
}
