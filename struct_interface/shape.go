package structinterface

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Length + rectangle.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (c Circle) Area() float64 {
	return 3.14 * math.Pow(c.Radius, 2)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
