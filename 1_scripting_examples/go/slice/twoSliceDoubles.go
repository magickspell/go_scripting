package main

import "fmt"

func intersection(slice1, slice2 []int) []int {
	// Создаем карту для хранения уникальных элементов первого слайса
	uniqueElements := make(map[int]bool)

	// Заполняем карту элементами первого слайса
	for _, num := range slice1 {
		uniqueElements[num] = true
	}

	// Создаем слайс для хранения пересечения
	var result []int

	// Проверяем элементы второго слайса на наличие в карте уникальных элементов
	for _, num := range slice2 {
		if uniqueElements[num] {
			// Если элемент присутствует в карте, добавляем его в результат
			result = append(result, num)
			// Удаляем элемент из карты, чтобы избежать повторного добавления
			delete(uniqueElements, num)
		}
	}

	return result
}

func main() {
	// Пример использования функции intersection
	slice1 := []int{1, 2, 3, 4, 5, 7}
	slice2 := []int{3, 4, 5, 6, 7, 8}

	fmt.Println("Intersection:", intersection(slice1, slice2)) // Вывод: [3 4 5]
}
