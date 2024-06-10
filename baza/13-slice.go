package main

import (
	"fmt"
)

func main() {
	// ptr = <array>
	// len 2
	// cap = 5
	array := [5]int64{1, 2, 3, 4, 5}

	slice := array[0:2] // [1connectionSimple:3] от до не включая ([2,3] cap = 4), все [:], от [1connectionSimple:], до [:2]

	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println("len(slice)", len(slice))
	fmt.Println("cap(slice)", cap(slice))

	slice = append(slice, 1)
	fmt.Println(slice)

	slice2 := []string{"a", "b", "c"}
	fmt.Println(slice2)
	t := 10
	slice3 := make([]int, t) // len = 10  cap = 10
	fmt.Println(slice3)
	slice4 := make([]int, 5, t) // len = 5  cap = 10
	fmt.Println(slice4)
	var nilSlice []int // тут nil значения но GO покажет как пустой массив
	fmt.Println(nilSlice)
	var sliceOfSlice = [][]int{} // слайс из слайсов
	fmt.Println(sliceOfSlice)
	/*
		если два слайса смотрят на один массив - при опреациях они могут повредить изначальные данные т.к. слайс ссылается на область памяти
		если нужно, то надо выделять новые слайсы вот так [:3:3] - т.е. нужно указывать третий максимальный элемент
	*/
}
