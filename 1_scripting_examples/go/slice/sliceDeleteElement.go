package main

import "fmt"

func main() {
	// Исходный слайс
	numbers := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	index := 2

	// Создание нового слайса без удаляемого элемента
	numbers = append(numbers[:index], numbers[index+1:]...)

	fmt.Println("Numbers:", numbers) // Output: [1 2 4 5]
}
