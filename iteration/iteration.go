package iteration

// Repeat takes a string and transforms it into
// the same string repeated multiple times (defined by "n").
func Repeat(s string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += s
	}
	return
}
