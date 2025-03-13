package main

import "fmt"

// интерфес состоит из 2 указателей, на тип и на данные
type MyErrorI interface {
	Error() string
}

type MyError struct {
	data string
}

func (e MyError) Error() string {
	return e.data
}

func main() {
	err := foo(4)
	// указатель на тип данных не nil
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok")
	}
}

func foo(i int) error {
	var err *MyError
	if i > 5 {
		err = &MyError{data: "i > 5"}
	}
	return err
}
