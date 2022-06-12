package countdown_test

import (
	"bytes"
	"hello/mocking/countdown"
	"testing"
)

type SpySleeper struct {
	CallsCounter int
}

func (s *SpySleeper) Sleep() {
	s.CallsCounter++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	countdown.Countdown(buffer, spySleeper)

	want := `3
2
1
Go!`
	got := buffer.String()

	if want != got {
		t.Errorf("want %q got %q", want, got)
	}

	if spySleeper.CallsCounter != 3 {
		t.Errorf("not enough calls to sleeper. Want 3 got %d", spySleeper.CallsCounter)
	}
}
