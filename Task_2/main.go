package main

import (
	"fmt"
	"series/fibonacci"
)

func main() {

	var number int
	defer fmt.Println("\nsuccessfully withdrawn!")
	fmt.Println("To get the fibonacci, enter their number:")
	fmt.Scan(&number)
	fibonacci.PrintLoop(number)
}
