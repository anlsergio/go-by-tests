package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	want := 4
	got := Add(2, 2)

	if got != want {
		t.Errorf("want '%d' but got '%d'", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
