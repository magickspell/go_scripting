package main

import (
	"fmt"
)

// customAppend добавляет элементы к слайсу и возвращает новый слайс с добавленными элементами.
func customAppend(slice []int, elements ...int) []int {
	// Вычисляем общую необходимую длину нового слайса
	totalLength := len(slice) + len(elements)
	// Проверяем, достаточно ли вместимости в исходном слайсе
	if totalLength <= cap(slice) {
		// Достаточно места, расширяем слайс
		slice = slice[:totalLength]
	} else {
		// Недостаточно места, создаем новый слайс с необходимой вместимостью
		newSlice := make([]int, totalLength)
		copy(newSlice, slice)
		slice = newSlice
	}
	// Копируем элементы в конец слайса
	copy(slice[len(slice)-len(elements):], elements)
	return slice
}

func main() {
	originalSlice := []int{1, 2, 3}
	// Добавление элементов в слайс
	extendedSlice := customAppend(originalSlice, 4, 5, 6)
	fmt.Println("Original Slice:", originalSlice)
	fmt.Println("Extended Slice:", extendedSlice)
}
