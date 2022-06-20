package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
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
			got := secondsInRadians(c.time)

			if c.angle != got {
				t.Fatalf("want %v radians, got %v", c.angle, got)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{
			time:  simpleTime(0, 0, 30),
			point: Point{0, -1},
		},
		{
			time:  simpleTime(0, 0, 45),
			point: Point{-1, 0},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)

			if !roughlyEqualPoint(c.point, got) {
				t.Fatalf("want %v point, got %v", c.point, got)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7

	return math.Abs(a-b) < equalityThreshold
}
