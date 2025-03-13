package main

import "fmt"

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func DoSomething() error {
	return &MyError{"Something went wrong НОВАЯ МОЯ ОШИБОЧКА ОШИБКА!"}
}

func main() {
	err := DoSomething()
	if err != nil {
		switch e := err.(type) {
		case *MyError:
			fmt.Printf("Custom error occurred: %s\n", e.Error())
		default:
			fmt.Println("Generic error occurred")
		}
	}
}
