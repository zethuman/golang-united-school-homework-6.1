package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

// CalcPerimeter returns calculation result of perimeter
func (t Triangle) CalcPerimeter() float64 {
	return t.Side + t.Side + t.Side
}

// CalcArea returns calculation result of area
func (t Triangle) CalcArea() float64 {
	semi := t.CalcPerimeter() / 2
	return math.Sqrt(semi * (semi - t.Side) * (semi - t.Side) * (semi - t.Side))
}
