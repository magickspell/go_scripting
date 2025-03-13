package main

import (
	"fmt"
	// "golang.org/x/exp/constraints"
)

func Max[T comparable](slice []T) T {
	// func Max[T constraints.Ordered](slice []T) T {
	max := slice[0]
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	ints := []int{1, 3, 2, 5, 4}
	fmt.Println("Max int:", Max(ints))

	floats := []float64{3.14, 1.618, 2.718, 2.0}
	fmt.Println("Max float64:", Max(floats))
}
