package main

import (
	"fmt"
)

func main() {
	str := "abc"
	for key, value := range str {
		fmt.Println(key, value)
	}
	fmt.Println(str[0])
	fmt.Println("---")
	strSlice := []byte("efg")
	for key, value := range strSlice {
		fmt.Println(key, value)
	}
	fmt.Println(strSlice[0])
	fmt.Println("---")
	strSliceRune := []rune("hij")
	for key, value := range strSliceRune {
		fmt.Println(key, value)
	}
	fmt.Println(strSliceRune[0])
}
