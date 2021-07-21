package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func(c Circle) Area() float64{
	if c.Radius<0{
		return 0
	}
	return math.Pi*c.Radius *c.Radius
}

func(c Circle) Perimeter() float64 {
	if c.Radius<0{
		return 0
	}
	return 2*math.Pi*c.Radius
}

func (c Circle) String() string {
	if c.Radius<0{
		println("Error")
	}
	return fmt.Sprintf("Circle:radius %.2f",c.Radius)
}
