/*
Starting with this code, sort the []user by
	age
	last
Also sort each []string “Sayings” for each user
	print everything in a way that is pleasant
*/
package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByLast []user

func (la ByLast) Len() int           { return len(la) }
func (la ByLast) Less(i, j int) bool { return la[i].Last < la[j].Last }
func (la ByLast) Swap(i, j int)      { la[i], la[j] = la[j], la[i] }

func Info(users []user) {
	fmt.Println("-------------")
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Println("\t", s)
		}
		fmt.Println("-------------")
		fmt.Println()
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	fmt.Println("-------------")
	fmt.Println("Sorted by age")
	sort.Sort(ByAge(users))
	Info(users)

	fmt.Println("Sorted by last")
	sort.Sort(ByLast(users))
	Info(users)

	for _, user := range users {
		sort.Strings(user.Sayings)
	}
	fmt.Println("With sorted sayings")
	Info(users)

}
