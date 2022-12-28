package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func (t testing.TB, shape Shape, want float64)  {
		got := shape.Area()
		
		if got != want {
			t.Errorf("got %g wanted %g", got, want)
		}
	}

	t.Run("cal area of circle", func(t *testing.T) {
		circle := Circle{radius: 10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g wanted %g", got, want)
		}
	})

	t.Run("calc area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12., 6.}
		got := rectangle.Area()
		want := 72.
		if got != want {
			t.Errorf("got %.2f but wanted %.2f", got, want)
		}
	})

	t.Run("calc area of a square", func(t *testing.T) {
		square := Square{4.}
		got:= square.Area()
		want := 16.

		if got != want {
			t.Errorf("got %f wanted %f ", got, want)
		}
	})
}
