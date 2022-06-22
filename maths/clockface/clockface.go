package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength   = 90
	minuteHandLength   = 80
	hourHandLength     = 50
	clockCenterX       = 150
	clockCenterY       = 150
	secondsInHalfClock = 30
	minutesInHalfClock = 30
	fullTurnInMinutes  = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	fullTurnInHours    = 2 * hoursInHalfClock
	svgStart           = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

// Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SVGWriter writes an SVG representation of an analogue clock to the writer w,
// considering the time t.
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := buildHand(secondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := buildHand(minuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := buildHand(hourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		p.X, p.Y)
}

func buildHand(p Point, length float64) Point {
	p = scaleHandLength(p, length)
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

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / (float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	// for every second, the minute hand will move 1/60th of the angle the second hand moves
	angleBasedInSeconds := secondsInRadians(t) / fullTurnInMinutes
	// movement of the minute hand itself based on minutes only
	angleBasedInMinutes := math.Pi / (minutesInHalfClock / float64(t.Minute()))
	return angleBasedInSeconds + angleBasedInMinutes
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	angleBasedInMinutes := minutesInRadians(t) / fullTurnInHours
	angleBasedInHours := math.Pi / (hoursInHalfClock / float64(convert24To12HourClockFormat(t)))
	return angleBasedInMinutes + angleBasedInHours
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func convert24To12HourClockFormat(t time.Time) int {
	return t.Hour() % fullTurnInHours
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{
		X: x,
		Y: y,
	}
}
