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

func TestCollectionAdd(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		want := 15
		got := CollectionAdd(numbers)

		if want != got {
			t.Errorf("want '%d', got '%d', given %v", want, got, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		want := 21
		got := CollectionAdd(numbers)

		if want != got {
			t.Errorf("want '%d', got '%d', given %v", want, got, numbers)
		}
	})
}
