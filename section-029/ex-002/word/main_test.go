package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	x := Count("Hello, world. How's it going?")
	if x != 5 {
		t.Error("Expected", 5, "got", x)
	}
}

func ExampleCount() {
	fmt.Println(Count("Hello, world. How's it going?"))
	// Output:
	// 5
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("Hello, world. How's it going?")
	}
}

func TestUseCount(t *testing.T) {
	x := UseCount("Hello Hello world. How's it going?")
	expected := map[string]int{
		"Hello":  2,
		"world.": 1,
		"How's":  1,
		"it":     1,
		"going?": 1,
	}
	if len(x) != len(expected) {
		t.Error("Expected", expected, "got", x)
	}
	for k, v := range expected {
		if x[k] != v {
			t.Error("Expected", expected, "got", x)
		}
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("Hello Hello world. How's it going?"))
	// Output:
	// map[Hello:2 How's:1 going?:1 it:1 world.:1]
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("Hello Hello world. How's it going?")
	}
}
