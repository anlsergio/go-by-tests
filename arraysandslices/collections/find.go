package collections

func Find[T any](items []T, predicate func(T) bool) (value T, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}

	return
}
