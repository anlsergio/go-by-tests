package _select

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://www.facebook.com"
	fastURL := "https://www.google.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
