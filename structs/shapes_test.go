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
		shape Shaper
		want  float64
	}{
		{
			shape: Rectangle{
				Width:  12.0,
				Height: 6.0,
			},
			want: 72.0,
		},
		{
			shape: Circle{
				Radius: 10,
			},
			want: 314.1592653589793,
		},
	}

	for _, tt := range testCases {
		got := tt.shape.Area()

		if tt.want != got {
			t.Errorf("want %g got %g", tt.want, got)
		}
	}
}
