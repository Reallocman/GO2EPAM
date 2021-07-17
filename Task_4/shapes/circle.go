package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func(c Circle) Area() float64{
	return math.Pi*c.Radius *c.Radius
}

func(c Circle) Perimeter() float64 {
	return 2*math.Pi*c.Radius
}

func (c Circle) String() string{
	return fmt.Sprintf("Circle:radius %.2f",c.Radius)
}
