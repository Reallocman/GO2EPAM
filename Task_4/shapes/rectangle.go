package shapes

import "fmt"

type Rectangle struct {
	Height float64
	Width  float64
}

func(r Rectangle) Area() float64{
	if r.Height<0 || r.Width<0{
		return 0
	}
	return r.Height *r.Width
}

func(r Rectangle) Perimeter() float64{
	if r.Height<0 || r.Width<0{
		return 0
	}
	return 2*(r.Width +r.Height)
}

func (r Rectangle) String() string{
	if r.Height<0 || r.Width<0 {
		println("Error")
	}
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f",r.Height,r.Width)
}
