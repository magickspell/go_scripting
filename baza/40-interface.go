package main

import "fmt"

// интерфейс
type SomeInterface interface {
	SomeMethod(in string) (out string, err error)
}

// реализация интерфейса - это конкретный тип (например структура, но можно и примитивные типы в алиасе)
type SomeStruct struct {
	prefix string
}

func (s SomeStruct) SomeMethod(in string) (out string) {
	return s.prefix + in
}

func main() {
	var testInterface SomeInterface
	fmt.Println("value %#+v\n", testInterface)
	fmt.Println("pointer %#+v\n", &testInterface)

	var testInterface2 = SomeStruct{"I say: "}
	fmt.Println("value %#+v\n", testInterface2)
	fmt.Println("pointer %#+v\n", &testInterface2)
	fmt.Println("method %q\n", testInterface2.SomeMethod("i am interface method"))
}
