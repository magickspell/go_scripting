package main

import "fmt"

func main() {
	var val interface{} = 42
	fmt.Println("val is:", val)
	// Пример безопасного type assertion
	if num, ok := val.(int); ok {
		fmt.Println("Value is an integer:", num)
	} else {
		fmt.Println("Value is not an integer")
	}

	// Пример не безопасного type assertion
	// str := val.(string) // Вызовет panic, так как тип значения не соответствует ожидаемому
	// fmt.Println(str)
}
