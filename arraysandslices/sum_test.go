package arraysandslices_test

import (
	collections "hello/arraysandslices"
	"reflect"
	"testing"
)

func TestCollectionAdd(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		want := 21
		got := collections.Add(numbers)

		if want != got {
			t.Errorf("want '%d', got '%d', given %v", want, got, numbers)
		}
	})
}

func TestCollectionsAdd(t *testing.T) {
	want := []int{3, 9}
	got := collections.AddCollections([]int{1, 2}, []int{0, 9})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want '%v', got '%v'", want, got)
	}
}

func TestCollectionsTailAdd(t *testing.T) {
	assertSums := func(t testing.TB, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want '%v', got '%v'", want, got)
		}
	}

	t.Run("sum of populated slices", func(t *testing.T) {
		want := []int{5, 9}
		got := collections.AddTails([]int{1, 2, 3}, []int{0, 9})

		assertSums(t, want, got)
	})

	t.Run("ensure safety of empty slices", func(t *testing.T) {
		want := []int{0, 9}
		got := collections.AddTails([]int{}, []int{0, 9})

		assertSums(t, want, got)
	})
}
