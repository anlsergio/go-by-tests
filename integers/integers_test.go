package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 4
	got := Add(2, 2)

	if want != got {
		t.Errorf("want '%d' but got '%d'", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestArrayAdder(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	want := 15
	got := ArrayAdd(numbers)

	if want != got {
		t.Errorf("want '%d', got '%d', given %v", want, got, numbers)
	}
}
