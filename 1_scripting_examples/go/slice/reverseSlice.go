package main

import "fmt"

func reverse(slice []int) {
	// 1 2 3 4 5  ->  5 2 3 4 1  ->  5 4 3 2 1
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	// Пример использования функции reverse
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Before reverse:", numbers)
	reverse(numbers)
	fmt.Println("After reverse:", numbers)
}
