package structsmethodsandinterfaces

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
