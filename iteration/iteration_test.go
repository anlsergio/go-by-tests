package iteration

import "testing"

func TestIteration(t *testing.T) {
	want := "aaaaa"
	got := Repeat("a")

	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}
}
