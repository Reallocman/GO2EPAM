package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func(c Circle) Area() float64{
	return math.Pi*c.Radius *c.Radius
}
func(r Rectangle) Area() float64{
	return r.Height *r.Width
}
func(c Circle) Perimeter() float64 {
	return 2*math.Pi*c.Radius
}
func(r Rectangle) Perimeter() float64{
	return 2*(r.Width +r.Height)
}
func (c Circle) String() string{
	return fmt.Sprintf("Circle:radius %.2f",c.Radius)
}
func (r Rectangle) String() string{
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f",r.Height,r.Width)
}
