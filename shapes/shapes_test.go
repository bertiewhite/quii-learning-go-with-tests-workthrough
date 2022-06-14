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

func TestAreas(t *testing.T) {
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{6, 12}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{3, 4}, 6.0},
	}

	for _, tt := range areaTests {
		CheckArea(t, tt.shape, tt.expected)
	}
}

func TestPerims(t *testing.T) {
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{6, 12}, 36.0},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range areaTests {
		CheckPerim(t, tt.shape, tt.expected)
	}
}
