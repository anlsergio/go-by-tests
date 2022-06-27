package collections_test

import (
	"hello/arraysandslices/collections"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		input := []int{1, 2, 3}

		multiply := func(a, b int) int {
			return a * b
		}

		const want = 6
		AssertEqual(t, want, collections.Reduce(input, multiply, 1))
	})

	t.Run("concatenate strings together", func(t *testing.T) {
		input := []string{"a", "b", "c"}

		concatenate := func(a, b string) string {
			return a + b
		}

		const want = "abc"
		AssertEqual(t, want, collections.Reduce(input, concatenate, ""))
	})
}

func AssertEqual[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Error("want true, got ", got)
	}
}
