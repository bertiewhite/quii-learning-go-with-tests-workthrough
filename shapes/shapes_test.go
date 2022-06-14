package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckArea(t testing.TB, shape Shape, expectation float64) {
	t.Helper()
	area := shape.Area()
	assert.Equal(t, area, expectation)
}

func CheckPerim(t testing.TB, shape Shape, expectation float64) {
	t.Helper()
	perim := shape.Perimeter()
	assert.Equal(t, perim, expectation)
}

func TestRectArea(t *testing.T) {
	t.Run("Find the area of a rectangle", func(t *testing.T) {
		rect := Rectangle{
			Width:  10.0,
			Height: 20.0,
		}
		expArea := 200.0

		CheckArea(t, &rect, expArea)
	})
}

func TestRectPerimeter(t *testing.T) {
	t.Run("Test perimeter of rectangle", func(t *testing.T) {
		rect := Rectangle{
			Width:  10.0,
			Height: 20.0,
		}
		expPerimeter := 60.0

		CheckPerim(t, &rect, expPerimeter)
	})
}

func TestCircleArea(t *testing.T) {
	t.Run("Calculates area of the circle", func(t *testing.T) {
		circle := Circle{10.0}
		expected := 314.1592653589793

		CheckArea(t, &circle, expected)
	})
}

func TestCirclePerimeter(t *testing.T) {
	t.Run("Calculates perimeter of the circle", func(t *testing.T) {
		circle := Circle{10.0}
		expected := 62.83185307179586

		CheckPerim(t, &circle, expected)
	})
}
