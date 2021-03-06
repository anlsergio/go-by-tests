package clockface_test

import (
	"hello/maths/clockface"
	"math"
	"testing"
	"time"
)

type radianCase struct {
	time  time.Time
	angle float64
}

type pointCase struct {
	time  time.Time
	point clockface.Point
}

func TestSecondsInRadians(t *testing.T) {
	cases := []radianCase{
		{
			time:  simpleTime(0, 0, 30),
			angle: math.Pi,
		},
		{
			time:  simpleTime(0, 0, 0),
			angle: 0,
		},
		{
			time:  simpleTime(0, 0, 45),
			angle: (math.Pi / 2) * 3,
		},
		{
			time:  simpleTime(0, 0, 7),
			angle: (math.Pi / 30) * 7,
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondsInRadians(c.time)
			assertRadians(t, c.angle, got)
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []pointCase{
		{
			time:  simpleTime(0, 0, 30),
			point: clockface.Point{0, -1},
		},
		{
			time:  simpleTime(0, 0, 45),
			point: clockface.Point{-1, 0},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondHandPoint(c.time)
			assertPoint(t, c.point, got)
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []radianCase{
		{
			time:  simpleTime(0, 30, 0),
			angle: math.Pi,
		},
		{
			time:  simpleTime(0, 0, 7),
			angle: 7 * (math.Pi / (30 * 60)),
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.MinutesInRadians(c.time)
			assertRadians(t, c.angle, got)
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []pointCase{
		{
			time:  simpleTime(0, 30, 0),
			point: clockface.Point{X: 0, Y: -1},
		},
		{
			time:  simpleTime(0, 45, 0),
			point: clockface.Point{X: -1, Y: 0},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.MinuteHandPoint(c.time)
			assertPoint(t, c.point, got)
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []radianCase{
		{
			time:  simpleTime(6, 0, 0),
			angle: math.Pi,
		},
		{
			time:  simpleTime(0, 0, 0),
			angle: 0,
		},
		{
			time:  simpleTime(21, 0, 0),
			angle: math.Pi * 1.5,
		},
		{
			time:  simpleTime(0, 1, 30),
			angle: math.Pi / ((6 * 60 * 60) / 90),
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.HoursInRadians(c.time)
			assertRadians(t, c.angle, got)
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []pointCase{
		{
			time:  simpleTime(6, 0, 0),
			point: clockface.Point{0, -1},
		},
		{
			time:  simpleTime(21, 0, 0),
			point: clockface.Point{-1, 0},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.HourHandPoint(c.time)
			assertPoint(t, c.point, got)
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func assertRadians(t *testing.T, want float64, got float64) {
	if !roughlyEqualFloat64(want, got) {
		t.Fatalf("want angle %v, got %v", want, got)
	}
}

func assertPoint(t *testing.T, want clockface.Point, got clockface.Point) {
	if !roughlyEqualPoint(want, got) {
		t.Fatalf("want %v point, got %v", want, got)
	}
}

func roughlyEqualPoint(a, b clockface.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7

	return math.Abs(a-b) < equalityThreshold
}
