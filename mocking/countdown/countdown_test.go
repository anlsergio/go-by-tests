package countdown_test

import (
	"bytes"
	"hello/mocking/countdown"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	countdown.Countdown(buffer)

	want := `3
2
1
Go!`
	got := buffer.String()

	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
