package main

import "fmt"

func main() {
	// Создание слайса с помощью литерала слайса
	numbers := []int{1, 2, 3, 4, 5}

	// Добавление элемента в конец слайса
	numbers = append(numbers, 6)

	// Изменение значения элемента слайса
	numbers[0] = 10

	// Получение длины и емкости слайса
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	// Вывод элементов слайса
	fmt.Println("Numbers:", numbers)

	// увеличть размер слайса на 10 элементов, не добавляя елементов
	numbers = append(numbers, make([]int, 10)...)
}
