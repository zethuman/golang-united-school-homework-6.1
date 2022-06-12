package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

var (
	errorShapeDoesNotExist   = errors.New("shape by index doesn't exist")
	errorIndexOutOfRange     = errors.New("index went out of the range")
	errorShapeCapacityOut    = errors.New("shapeCapacity range out")
	errorCirclesDoesNotExist = errors.New("circles are not exist in the list")
)

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == 0 {
		b.shapes = make([]Shape, 0)
	}

	if len(b.shapes) >= b.shapesCapacity {
		return errorShapeCapacityOut
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errorIndexOutOfRange
	}

	if _, ok := b.shapes[i].(Shape); !ok {
		return nil, errorShapeDoesNotExist
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errorIndexOutOfRange
	}

	removed := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return removed, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errorIndexOutOfRange
	}

	removed := b.shapes[i]
	replaced := append(b.shapes[:i], shape)
	b.shapes = append(replaced, b.shapes[i+1:]...)

	return removed, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i += 1 {
		sum += b.shapes[i].CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var area float64
	for i := 0; i < len(b.shapes); i += 1 {
		area += b.shapes[i].CalcArea()
	}
	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var counter, i int

	fmt.Println("Length before: ", len(b.shapes))

	for len(b.shapes) > i {
		if _, ok := b.shapes[i].(Circle); !ok {
			i++
			continue
		}
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		counter++
	}

	if counter == 0 {
		return errorCirclesDoesNotExist
	}

	fmt.Println("Length after: ", len(b.shapes))
	fmt.Println(b.shapes)

	return nil
}
