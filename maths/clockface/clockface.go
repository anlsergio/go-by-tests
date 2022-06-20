package clockface

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCenterX     = 150
	clockCenterY     = 150
)

// Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an
// analogic clock positioned at time `t` represented as a Point
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = scaleHandLength(p, secondHandLength)
	p = flipOverTheXAxis(p)
	p = translateToTheCenterPosition(p)

	return p
}

func translateToTheCenterPosition(p Point) Point {
	return Point{
		X: p.X + clockCenterX,
		Y: p.Y + clockCenterY,
	}
}

// flip it over the X axis because of the SVG origin point
func flipOverTheXAxis(p Point) Point {
	return Point{
		X: p.X,
		Y: -p.Y,
	}
}

func scaleHandLength(p Point, length float64) Point {
	return Point{
		X: length * p.X,
		Y: length * p.Y,
	}
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / (float64(t.Second())))
}
