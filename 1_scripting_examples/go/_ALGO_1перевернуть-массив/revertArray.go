package main

import (
	"fmt"
	"math"
)

func processArray(arr []int) []int {
	lnt := len(arr)
	midFloat := float32(lnt) / 2
	mid := int(math.Round(float64(midFloat)))
	fmt.Println("[midFloat] = ", midFloat)
	fmt.Println("[mid] = ", mid)

	result := make([]int, 0)

	for idx := 0; idx < mid; idx++ {
		result = append(result, arr[lnt-idx-1])
	}

	past := arr[0 : mid-1]
	fmt.Println("past = ", past)
	pastLen := len(past) - 1
	for ; pastLen >= 0; pastLen-- {
		result = append(result, past[pastLen])
	}

	return result
}

func processArray2(arr []int) []int {
	lnt := len(arr)

	result := make([]int, lnt)

	for idx := 0; idx < lnt-1; idx++ {
		result[idx] = arr[lnt-idx-1]
		result[lnt-idx-1] = arr[idx]
	}

	return result
}

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	array2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	array3 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Printf("%v\n", processArray2(array1))
	fmt.Printf("%v\n", processArray2(array2))
	fmt.Printf("%v\n", processArray2(array3))
}
