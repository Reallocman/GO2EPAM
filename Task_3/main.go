package main

import (
	"fmt"
	"storages/average"
	"storages/maps"
	"storages/slices"
)
func main() {

	array:= [6]float64{1,2,3,4,5,6}
	fmt.Println(average.Mean(array))

	s1 := []string{"one","two","three"}
	fmt.Println(slices.Max(s1))

	s2 := []int64{1,2,5,15}
	fmt.Println(slices.Reverse(s2))

	m:= map[int]string{7:"kto",4:"pochemy"}
	maps.PrintSorted(m)
}

