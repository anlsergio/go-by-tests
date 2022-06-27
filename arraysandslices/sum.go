package arraysandslices

import "hello/arraysandslices/collections"

// Add returns the total sum of the numbers being passed in.
func Add(numbers []int) int {
	addAccumulator := func(a, b int) int {
		return a + b
	}
	return collections.Reduce(numbers, addAccumulator, 0)
}

// AddCollections returns the sum of each individual collection being passed in
func AddCollections(c ...[]int) []int {
	addCollectionsAccumulator := func(a, b []int) []int {
		return append(a, Add(b))
	}

	return collections.Reduce(c, addCollectionsAccumulator, []int{})
}

// AddTails returns the tail sum of each individual collection being passed in
func AddTails(c ...[]int) []int {
	addTailAccumulator := func(a, b []int) []int {
		if len(b) == 0 {
			return append(a, 0)
		}
		tail := b[1:]
		return append(a, Add(tail))
	}

	return collections.Reduce(c, addTailAccumulator, []int{})
}
