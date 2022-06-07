package integers

// Add takes two integers and returns their sum result.
func Add(n1, n2 int) int {
	return n1 + n2
}

// ArrayAdd returns the total sum of the numbers being passed in.
func ArrayAdd(numbers [5]int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}
