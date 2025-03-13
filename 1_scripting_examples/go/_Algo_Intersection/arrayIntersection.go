package main

import "fmt"

func intersection(one, two []int) []int {
	var m = make(map[int]uint) // не делаем пре-аллокацию, так как не знаем количество дублей

	for i := range one { // пробегаясь по первому слайсу "прогреваем" карту
		m[one[i]]++ // так как нулевое значение для uint это 0 - то просто увеличиваем
	}

	var result = make([]int, 0) // тоже без пре-аллокации, т.к. не знаем сколько пересечений

	for i := range two { // пробегаясь по второму - ищем пересечение
		if value, ok := m[two[i]]; ok {
			if value > 0 {
				m[two[i]]--
				result = append(result, two[i])
			} else {
				delete(m, two[i]) // прибираемся, так как ключ уже не нужен (== 0)
			}
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", intersection([]int{23, 3, 1, 2}, []int{6, 2, 4, 23})) // [2, 23]
	fmt.Printf("%v\n", intersection([]int{1, 1, 1}, []int{1, 1, 1, 1}))      // [1, 1, 1]
}
