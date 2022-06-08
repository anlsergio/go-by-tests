package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	want := 40.0
	got := Perimeter(rectangle)

	if want != got {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		name    string
		shape   Shaper
		hasArea float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				Width:  12,
				Height: 6,
			},
			hasArea: 72.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Radius: 10,
			},
			hasArea: 314.1592653589793,
		},
		{
			name: "Triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			hasArea: 36.0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if tt.hasArea != got {
				t.Errorf("%#v want %g got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}
