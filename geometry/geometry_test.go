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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g wanted %g", got, want)
		}
	}

	t.Run("cal area of circle", func(t *testing.T) {
		circle := Circle{radius: 10}
		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("calc area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12., 6.}

		checkArea(t, rectangle, 72)
	})

	t.Run("calc area of a square", func(t *testing.T) {
		square := Square{4.}
		checkArea(t,square,16)
	})
}
