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
	rectangle := Rectangle{
		Width:  12.0,
		Height: 6.0,
	}
	want := 72.0
	got := Area(rectangle)

	if want != got {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}
