package clockface_test

import (
	"bytes"
	"encoding/xml"
	"hello/maths/clockface"
	"testing"
	"time"
)

//func TestSecondHandAtMidnight(t *testing.T) {
//	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
//
//	var centerAxis float64 = 150
//
//	want := clockface.Point{
//		X: centerAxis,
//		Y: centerAxis - 90,
//	}
//	got := clockface.SecondHand(tm)
//
//	if want != got {
//		t.Errorf("Want %v, got %v", want, got)
//	}
//}
//
//func TestSecondHandAt30Seconds(t *testing.T) {
//	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
//
//	var centerAxis float64 = 150
//
//	want := clockface.Point{
//		X: centerAxis,
//		Y: centerAxis + 90,
//	}
//	got := clockface.SecondHand(tm)
//
//	if want != got {
//		t.Errorf("Want %v, got %v", want, got)
//	}
//}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Lines   []Line   `xml:"line"`
}

func TestSVGWriterAtMidnight(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			time: simpleTime(0, 0, 0),
			line: Line{X1: 150, Y1: 150, X2: 150, Y2: 60},
		},
		{
			time: simpleTime(0, 0, 30),
			line: Line{X1: 150, Y1: 150, X2: 150, Y2: 240},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Lines) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Lines)
			}
		})
	}

}

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if l == line {
			return true
		}
	}

	return false
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
