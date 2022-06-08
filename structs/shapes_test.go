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
	t.Run("rectangle calculation", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  12.0,
			Height: 6.0,
		}
		want := 72.0
		got := rectangle.Area()

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("circle calculation", func(t *testing.T) {
		circle := Circle{
			Radius: 10,
		}
		want := 314.1592653589793
		got := circle.Area()

		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})
}
