package main

import (
	"fmt"
)

func main() {
	var arr [3]float64 = [3]float64{1, 2, 3.14}
	fmt.Println(arr)

	arr2 := [2]int64{1, 2}
	fmt.Println(arr2)

	empty := []string{}
	fmt.Println(empty)
}
