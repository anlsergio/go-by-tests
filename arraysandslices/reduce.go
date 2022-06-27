package collections

func Reduce[T any](iterable []T, accumulator func(T, T) T, initialValue T) T {
	// The initial value is used as "The Identity Element":
	// 	In mathematics, an identity element, or neutral element,
	// 	of a binary operation operating on a set is an element of the set which
	//	leaves unchanged every element of the set when the operation is applied.
	// 	Examples:
	//		- Sum: 0 + 5 = 5 (the Identity Element is 0 for Sum operations)
	//		- Multiplication: 1 * 5 = 5 (the Identity Element is 1 for Multiplication operations)
	var result = initialValue
	for _, c := range iterable {
		result = accumulator(result, c)
	}

	return result
}
