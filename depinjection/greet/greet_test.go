package greet_test

import (
	"bytes"
	"hello/depinjection/greet"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	greet.Greet(&buffer, "Chris")

	want := "Hello, Chris"
	got := buffer.String()

	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
