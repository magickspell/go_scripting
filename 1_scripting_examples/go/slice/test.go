package main

import "fmt"

// Функция для добавления элемента
func addElement(slice []int, element int) []int {
	return append(slice, element)
}

// Функция фильтрации слайса
func filterSlice(slice []int, filterFunc func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if filterFunc(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Пустой слайс и слайс с предопределенными значениями
	var emptySlice []int
	fmt.Println("[emptySlice]")
	fmt.Println(len(emptySlice))
	fmt.Println(cap(emptySlice))

	predefinedSlice := []int{1, 2, 3}
	fmt.Println("[predefinedSlice]")
	fmt.Println(len(predefinedSlice))
	fmt.Println(cap(predefinedSlice))

	// Создание слайса с make
	sizedSlice := make([]int, 1, 5) // Слайс с емкостью 5 b и длинной 1
	fmt.Println("[sizedSlice]")
	fmt.Println(len(sizedSlice))
	fmt.Println(cap(sizedSlice))

	// Инициализация слайса с литералом
	literalSlice := []int{4, 5, 6, 7, 8}
	fmt.Println("[literalSlice]")
	fmt.Println(len(literalSlice))
	fmt.Println(cap(literalSlice))

	// Добавление элементов
	nums := append(predefinedSlice, 4)
	// Удаление элемента (сохранение порядка)
	nums = append(nums[:1], nums[2:]...)
	// Вставка элемента
	middleIndex := len(nums) / 2
	nums = append(nums[:middleIndex], append([]int{100}, nums[middleIndex:]...)...)

	numsTest := make([]int, 0, 1) // Начните с малой емкости
	for i := 0; i < 100; i++ {
		numsTest = append(numsTest, i)
		fmt.Println("Длина:", len(numsTest), "Емкость:", cap(numsTest))
	}

	boolTestFunc := func(i int) bool { return i > 95 }
	newNumsTestAfterFunc := filterSlice(numsTest, boolTestFunc)
	fmt.Println("[newNumsTestAfterFunc]")
	fmt.Println(newNumsTestAfterFunc)
}
