package shapes

import (
	"errors"
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() (float64,error) {
	if c.Radius < 0 {
		return 0,errors.New("you have entered incorrect parameters")
	}
	return math.Pi * c.Radius * c.Radius,nil
}

func (c Circle) Perimeter() (float64,error) {
	if c.Radius < 0 {
		return 0,errors.New("you have entered incorrect parameters")
	}
	return 2 * math.Pi * c.Radius,nil
}

func (c Circle) String() string {
	if c.Radius < 0 {
		return fmt.Sprintf("Circle:radius %.2f", c.Radius)
	}
	return fmt.Sprintf("Circle:radius %.2f", c.Radius)
}
