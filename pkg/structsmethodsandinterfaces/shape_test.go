package structsmethodsandinterfaces

import (
	"math"
	"testing"
)

func TestShape(t *testing.T) {
	check := func(t testing.TB, expected, actual float64, shape Shape) {
		t.Helper()

		if actual != expected {
			t.Errorf("Expected: %g - Actual: %g - Shape: %v", expected, actual, shape)
		}
	}

	checkArea := func(t *testing.T, expected float64, shape Shape) {
		actual := shape.Area()

		check(t, expected, actual, shape)
	}

	checkPerimeter := func(t *testing.T, expected float64, shape Shape) {
		actual := shape.Perimeter()

		check(t, expected, actual, shape)
	}

	t.Run("Circle", func(t *testing.T) {
		circle := &Circle{radius: 3.0}
		checkArea(t, 9.0*math.Pi, circle)
		checkPerimeter(t, 6.0*math.Pi, circle)
	})

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := &Rectangle{height: 3.0, width: 4.0}
		checkArea(t, 12.0, rectangle)
		checkPerimeter(t, 14.0, rectangle)
	})
}
