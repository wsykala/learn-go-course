/*
Start with this code. Get the code ready to BET on (add benchmarks, examples, tests) however do not write an example for the func that returns a map; and only write a test for it as an extra challenge.
Add documentation to the code. Run the following in this order:
	tests
	benchmarks
	coverage
	coverage shown in web browser
	examples shown in documentation in a web browser
*/

package main

import (
	"fmt"

	"github.com/wsykala/learn-go-course/section-029/ex-002/quote"
	"github.com/wsykala/learn-go-course/section-029/ex-002/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
