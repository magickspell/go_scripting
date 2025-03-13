package main

import "fmt"

func main() {
	// деферы вызываются снизу вверх (в обратном порядке их объявления, FILO)
	fmt.Println(0)
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(42)
}
