package depinjection_test

import (
	"bytes"
	"hello/depinjection"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	depinjection.Greet(&buffer, "Chris")

	want := "Hello, Chris"
	got := buffer.String()

	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
