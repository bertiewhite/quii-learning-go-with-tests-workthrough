package shapes

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Height * t.Base
}

// this basic understanding of triangles breaks down at this point: Isocolies/right angled/scalene triangles all require
// special treatment.
func (t Triangle) Perimeter() float64 {
	return 0
}
