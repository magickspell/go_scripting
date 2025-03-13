package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	myString := "1,2,3"
	myIntSlice := []int{}

	for idx, value := range strings.Split(myString, ",") {
		fmt.Println("idx: %v", idx)
		newInt, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("idx: %v", err)
		}
		myIntSlice = append(myIntSlice, newInt)
	}
	fmt.Println("myIntSlice: %v", myIntSlice)

	slice := []int{1, 2, 3, 4, 5, 1, 3}
	newMap := make(map[int]int)
	for _, value := range slice {
		val, ok := newMap[value]
		if !ok {
			fmt.Println("[val]")
			fmt.Println(val)
			newMap[value] = value
		}
	}
	fmt.Println("[newMap]")
	fmt.Println(newMap)
}
