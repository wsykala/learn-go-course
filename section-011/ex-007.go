/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.
*/
package main

import "fmt"

func main() {
	arr := make([][]string, 0, 2)
	arr = append(arr, []string{"James", "Bond", "Shaken, not stirred"})
	arr = append(arr, []string{"Miss", "Moneypenny", "Helloooooo, James."})

	for _, v := range arr {
		fmt.Printf("%s, %d\n", v, len(v))
		for _, s := range v {
			fmt.Printf("\t%s, %d\n", s, len(s))
		}
	}
}
