package iteration

// Repeat takes a string and transforms it into
// the same string repeated 5 times.
func Repeat(s string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += s
	}
	return
}
