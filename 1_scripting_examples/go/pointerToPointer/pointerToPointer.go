package main

import "fmt"

func main() {
	var a int = 42
	var ptr *int = &a       // Указатель на переменную 'a'
	var ptrptr **int = &ptr // Указатель на указатель 'ptr'

	fmt.Println("Значение переменной 'a':", a)
	fmt.Println("Значение, на которое указывает 'ptr':", *ptr)
	fmt.Println("Значение, на которое указывает 'ptrptr':", **ptrptr)

	// Изменение значения переменной 'a' через указатель на указатель
	**ptrptr = 100
	fmt.Println("Новое значение переменной 'a':", a)
}
