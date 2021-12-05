package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}
	cases := []test{
		{data: []int{1, 4, 6, 8, 100}, answer: 6},
		{data: []int{0, 8, 10, 1000}, answer: 9},
		{data: []int{9000, 4, 10, 8, 6, 12}, answer: 9},
		{data: []int{123, 744, 140, 200}, answer: 170},
	}
	for _, v := range cases {
		x := CenteredAvg(v.data)
		if v.answer != x {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}
