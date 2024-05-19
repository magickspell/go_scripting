package main

import (
	"fmt"
	"math"
)

// zero value у структу - структура заполненная нулевыми значениями
// гетеров и сетеров нет
// приватных, публичных, защищеных ключей нет
// НО - с маленькой буквы внутри пакета (папки) - доступно только в тьекущей папке, с БОЛЬШОЙ - везде
type Point struct { // новый тип на основе структуры Point
	X float64
	Y float64
	Z string
}

func (p Point) Abs() float64 { // так мы объявляем метод структуры (тут без указателя возвращается новая структура?)
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) SetX(newX float64) { // что бы менять туже структура на нее надо навесить указатель
	p.X = newX
}

func main() {
	p := Point{X: 1, Y: 37}
	fmt.Println(p)
	fmt.Println(p.Y)
	fmt.Printf("%v", p)
	fmt.Println("")
	fmt.Printf("%+v", p)
	fmt.Println("")
	fmt.Printf("%#+v", p)
	// незаполненные поля заполняются нулевыми значениями

	// сравнивать структуры можно только когда они состоят из примитивов
	fmt.Println("")
	fmt.Println(p == Point{X: 1})
	fmt.Println(p == Point{X: 1, Y: 37})

	// вызов метода струтуры (есть два способа)
	fmt.Println(p.Abs())
	fmt.Println(Point.Abs(p))

	// с указателем меняем переменную структуры
	p2 := Point{X: 1, Y: 2}
	fmt.Println(p2.X)
	p2.SetX(100)
	fmt.Println(p2.X)
}
