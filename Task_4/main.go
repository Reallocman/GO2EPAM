package main

import (
	"structures/interfaces"
)

func main() {

	c := interfaces.Circle{Radius: 8}
	r := interfaces.Rectangle{Height: 9, Width:3}
	interfaces.DescribeShape(c)
	interfaces.DescribeShape(r)

}
