package main

import "fmt"

func foo(src []int) {
	src = append(src, 5)
}

func main() {
	arr := []int{1, 2, 3}
	src := arr[:0]
	fmt.Println(src)

	foo(src)
	fmt.Println(src)
	fmt.Println(arr)
}
