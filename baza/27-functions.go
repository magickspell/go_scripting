package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(2, 2))
	fmt.Println(SumInfinit(2, 1, 2, 3, 4, 5))
	fmt.Println(SumInfinit(2, []int{1, 2, 8}...))

	// видимость i разная в функции и здесь
	i := 0
	inc(i)
	fmt.Println("")
	fmt.Print("main func i = ", i)

	// но мы можем успользовать указатель
	fmt.Println("")
	a := 0
	incPointer(&a)
	fmt.Println("")
	fmt.Print("main func a = ", a)
}

func sum(a int, b int) int { // доступна только внутри этого пакета - с МАЛЕНЬКОЙ буквы
	return a + b
}

func SumInfinit(i int, items ...int) (int, bool) { // доступна везде - с БОЛЬШОЙ букв
	result := i
	for _, value := range items {
		result += value
	}
	return result, true
}

func inc(i int) {
	i++
	fmt.Print("inc func i = ", i)
}

func incPointer(a *int) {
	*a++
	fmt.Print("inc func a = ", a)
}
