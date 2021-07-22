package shapes

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() (float64,error) {
	if r.Height < 0 || r.Width < 0 {
		return 0,errors.New("you have entered incorrect parameters")
	}
	return r.Height * r.Width,nil
}

func (r Rectangle) Perimeter() (float64,error) {
	if r.Height < 0 || r.Width < 0 {
		return 0,errors.New("you have entered incorrect parameters")
	}
	return 2 * (r.Width + r.Height),nil
}

func (r Rectangle) String() (string,error) {
	if r.Height < 0 || r.Width < 0 {
		return "",errors.New("you have entered incorrect parameters")
	}
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f", r.Height, r.Width),nil
}
