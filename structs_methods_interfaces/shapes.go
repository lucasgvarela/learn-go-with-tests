package shapes

import "math"

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Shape interface {
	Area() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculate the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculate the area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
