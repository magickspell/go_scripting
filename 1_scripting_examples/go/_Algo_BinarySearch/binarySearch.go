package main

import "fmt"

// binarySearch выполняет бинарный поиск элемента target в отсортированном массиве arr
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2 // Находим средний индекс
		// Проверяем, равен ли средний элемент искомому
		if arr[mid] == target {
			return mid
		}
		// Если искомый элемент меньше среднего, ищем в левой половине
		if arr[mid] > target {
			right = mid - 1
		} else {
			// Иначе ищем в правой половине
			left = mid + 1
		}
	}
	// Если элемент не найден, возвращаем -1
	return -1
}

func main() {
	arr := []int{2, 3, 4, 5, 6, 9, 10, 40, 41, 42, 1337}
	target := 10
	result := binarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Элемент найден на позиции %d\n", result)
	} else {
		fmt.Println("Элемент не найден в массиве")
	}
}
