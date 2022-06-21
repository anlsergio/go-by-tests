package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCenterX     = 150
	clockCenterY     = 150
	svgStart         = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
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
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)
	p = scaleHandLength(p, secondHandLength)
	p = flipOverTheXAxis(p)
	p = translateToTheCenterPosition(p)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		p.X, p.Y)
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
