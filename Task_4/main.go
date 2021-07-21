package main

import (
	"structures/shapes"
	)

func main() {

	c := shapes.Circle{Radius: 8}
	r := shapes.Rectangle{Height: 9, Width:3}
	shapes.DescribeShape(c)
	shapes.DescribeShape(r)

}
