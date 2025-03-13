package main

import "fmt"

// обертка
// Объявление нового типа, который является оберткой для строки
type MyString string

// Реализация метода для нового типа
func (ms MyString) Print() {
	fmt.Println(ms)
}

func main() {
	// Создание экземпляра нового типа
	myStr := MyString("Hello, world!")

	// Вызов метода для экземпляра нового типа
	myStr.Print()
}
