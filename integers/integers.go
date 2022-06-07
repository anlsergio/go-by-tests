package integers

// Add takes two integers and returns their sum result.
func Add(n1, n2 int) int {
	return n1 + n2
}

// CollectionAdd returns the total sum of the numbers being passed in.
func CollectionAdd(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

// CollectionsAdd returns the sum of each individual collection being passed in
func CollectionsAdd(collections ...[]int) []int {
	var sums []int

	for _, c := range collections {
		sums = append(sums, CollectionAdd(c))
	}

	return sums
}
