package main

import "fmt"

// Интерфейс для определения, может ли автомобиль ехать
type Drivable interface {
	Drive() string
}

// Интерфейс для определения, может ли автомобиль остановиться
type Stoppable interface {
	Stop() string
}

// Тип данных Car реализует оба интерфейса
type Car struct {
	Brand string
}

// Методы для реализации интерфейсов Drivable и Stoppable
func (c Car) Drive() string {
	return fmt.Sprintf("%s is driving", c.Brand)
}

func (c Car) Stop() string {
	return fmt.Sprintf("%s stopped", c.Brand)
}

func main() {
	myCar := Car{Brand: "Toyota"}
	fmt.Println(myCar.Drive()) // Вывод: Toyota is driving
	fmt.Println(myCar.Stop())  // Вывод: Toyota stopped
}
