package main

import (
	"fmt"
)

// BubbleSort сортирует массив чисел методом пузырьковой сортировки
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 1, 90}
	fmt.Println("Исходный массив:", arr)

	BubbleSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
