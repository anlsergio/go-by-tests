package collections_test

import (
	"hello/arraysandslices/collections"
	"strings"
	"testing"
)

type Person struct {
	Name string
}

func TestFind(t *testing.T) {
	t.Run("find even number", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		findEven := func(x int) bool {
			return x%2 == 0
		}

		firstEvenNumber, found := collections.Find(input, findEven)
		AssertTrue(t, found)
		AssertEqual(t, 2, firstEvenNumber)
	})

	t.Run("find the best programmer ever", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}

		findBestProgrammer := func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		}

		king, found := collections.Find(people, findBestProgrammer)

		AssertTrue(t, found)

		// I've got to agree with that.
		// Shout out to you Chris!!
		want := Person{
			Name: "Chris James",
		}
		AssertEqual(t, want, king)
	})
}
