package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// Метод для структуры Rectangle, вычисляющий площадь
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println(rect.Area()) // Вызов метода Area для экземпляра rect
}
