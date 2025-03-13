package main

import "fmt"

// Определение интерфейса
type Printer interface {
	Print()
}

// Конкретный тип, реализующий интерфейс
type MyStruct struct {
	Name string
}

// Реализация метода интерфейса
func (m MyStruct) Print() {
	fmt.Println("Name:", m.Name)
}

func main() {
	var x interface{} = MyStruct{Name: "Example"}

	// Проверка типа и реализации интерфейса
	switch v := x.(type) {
	case Printer:
		v.Print() // Безопасно вызывается, так как x реализует интерфейс Printer
	default:
		fmt.Println("x does not implement Printer")
	}
}
