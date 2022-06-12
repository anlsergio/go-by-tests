package countdown_test

import (
	"bytes"
	"hello/mocking/countdown"
	"reflect"
	"testing"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		countdown.Countdown(buffer, &SpyCountdownOperations{})

		want := `3
2
1
Go!`
		got := buffer.String()

		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		countdown.Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("want calls counter to be %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
