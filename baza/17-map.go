package main

import (
	"fmt"
)

func main() {
	mapStrNum := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(mapStrNum)
	fmt.Println(mapStrNum["two"])
	fmt.Println(mapStrNum["thousands"]) // нулевое значение
	value, ok := mapStrNum["thousands"] // проверка что значение из массива
	valueZ, okZ := mapStrNum["zero"]
	fmt.Println(value, ok)   // не значение
	fmt.Println(valueZ, okZ) // является значением
	// в ифе вот так удобно проверять
	// valueIf, okIf := mapStrNum["if"]
	// if ojIf
	if valueIf, okIf := mapStrNum["if"]; okIf { // из такой конструкции область видимости значений будет только во внутренних ветках
		fmt.Println("this is ", valueIf, true)
	} else {
		fmt.Println("this is ", valueIf, false)
	}

	mapNumStr := make(map[int64]string)
	fmt.Println(mapNumStr)
	mapNumStr[1] = "one"
	mapNumStr[2] = "two"
	fmt.Println(mapNumStr)
	mapNumStr[1] = "three"
	fmt.Println(mapNumStr)
}
