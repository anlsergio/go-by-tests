package iteration

import "testing"

func TestRepeat(t *testing.T) {
	want := "aaaaa"
	got := Repeat("a")

	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}
}
