package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	for index, value := range arr {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}
}
