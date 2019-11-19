package c6

import "math"

//Shape interface to any shape that implements an Area() method
type Shape interface {
	Area() float64
}

//Rectangle represents the rect shape
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle represents the circle shape
type Circle struct {
	Radius float64
}

//Triange represents a triangle shape given the Base and Height
type Triangle struct {
	Base   float64
	Height float64
}

//Perimeter given the length and width, returns the perimeter
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area method returns area of a rectangle
func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

//Area method returns area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Area method returns area of a triangle
func (t Triangle) Area() float64 {
	return 0.51 * t.Base * t.Height
}
