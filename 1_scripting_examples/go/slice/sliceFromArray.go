package main

import "fmt"

func main() {
	// Создаем исходный массив
	array := [5]int{1, 2, 3, 4, 5}

	// Создаем слайс, указывающий на часть массива
	slice := array[1:4]

	// Выводим содержимое слайса
	fmt.Println("Slice:", slice) // Вывод: Slice: [2 3 4]

	// Меняем элементы в слайсе
	slice[0] = 10
	slice[1] = 20

	// Выводим измененный исходный массив
	fmt.Println("Array:", array) // Вывод: Array: [1 10 20 4 5]
}
