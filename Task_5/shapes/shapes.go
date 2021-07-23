package shapes

import "fmt"

type Shape interface {
	Area() (float64,error)
	Perimeter() (float64,error)
}

func DescribeShape(s Shape) {
	area,err:=s.Area()
	if err!=nil {
		fmt.Println(err)
		return
	}
	perimetr,err:=s.Perimeter()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n",area)
	fmt.Printf("Perimeter: %.2f\n",perimetr)
}
