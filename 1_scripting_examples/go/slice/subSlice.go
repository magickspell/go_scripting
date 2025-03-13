package main

import "fmt"

func main() {
	// Исходный слайс
	numbers := []int{1, 2, 3, 4, 5}

	// Получение подслайса с индексами от 1 до 3 (не включительно)
	subSlice := numbers[1:3]
	fmt.Println("Subslice:", subSlice) // Output: [2 3]

	// Получение подслайса с индексами от 2 до конца слайса
	subSlice2 := numbers[2:]
	fmt.Println("Subslice2:", subSlice2) // Output: [3 4 5]

	// Получение подслайса с начала до индекса 3 (не включительно)
	subSlice3 := numbers[:3]
	fmt.Println("Subslice3:", subSlice3) // Output: [1 2 3]

	// Получение копии исходного слайса
	copySlice := numbers[:]
	fmt.Println("CopySlice:", copySlice) // Output: [1 2 3 4 5]
}
