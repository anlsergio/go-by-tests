package clockface_test

import (
	"hello/maths/clockface"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	var centerAxis float64 = 150

	want := clockface.Point{
		X: centerAxis,
		Y: centerAxis - 90,
	}
	got := clockface.SecondHand(tm)

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	var centerAxis float64 = 150

	want := clockface.Point{
		X: centerAxis,
		Y: centerAxis + 90,
	}
	got := clockface.SecondHand(tm)

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}
