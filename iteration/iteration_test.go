package iteration

import "testing"

func TestRepeat(t *testing.T) {
	want := "aaaaaaaa"
	got := Repeat("a", 8)

	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
