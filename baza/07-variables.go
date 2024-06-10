package main

import "fmt"

func main() {
	// var sum = 10
	// var sum int = 10
	var sum int // без значения, но тут будет 0 - т.к. записывается Zero Value
	// sum := 10 // покороче через оператор "моржа" )))

	var a rune   // int9, int16, int32. int64, byte, rune
	var b bool   // false, true (zero value = false)
	var s string // string (zero value = "")

	fmt.Println(sum)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s)

	const PI = 314
	const (
		A = 1
		// B string // так не может
		B         = "2"
		C bool    = true
		D float64 = 8 / 4
	)
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)

	fmt.Println("iota conts - побитово сдвигает и присваивает каждое новое значени в константу")
	const (
		RED = iota
		// RED = 1connectionSimple >> iota // 1connectionSimple 0 0
		// RED = 1connectionSimple << iota // 1connectionSimple 2 4
		// RED = 2 << iota // 2 4 8
		// RED = 3 << iota // 3 6 12
		GREEN
		BLUE
		_
		YELLOW
	)
	fmt.Println(RED)
	fmt.Println(GREEN)
	fmt.Println(BLUE)
	fmt.Println(YELLOW)
}
