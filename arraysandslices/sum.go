package collections

// Add returns the total sum of the numbers being passed in.
func Add(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

// AddCollections returns the sum of each individual collection being passed in
func AddCollections(collections ...[]int) []int {
	var sums []int

	for _, c := range collections {
		sums = append(sums, Add(c))
	}

	return sums
}

// TailAdd returns the tail sum of each individual collection being passed in
func TailAdd(collections ...[]int) []int {
	var sums []int

	for _, c := range collections {
		if len(c) == 0 {
			sums = append(sums, 0)
		} else {
			tail := c[1:]
			sums = append(sums, Add(tail))
		}
	}

	return sums
}
