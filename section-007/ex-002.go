/*
Using the following operators, write expressions and assign their values to variables:
	==
	<=
	>=
	!=
	<
	>
Now print each of the variables.
*/

package main

import "fmt"

func main() {
	g := (12 == 12)
	h := (12 <= 13)
	i := (12 >= 13)
	j := (12 != 12)
	k := (12 < 13)
	l := (12 > 12)
	fmt.Println(g, h, i, j, k, l)
}
