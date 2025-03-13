package main

import "fmt"

// для безопасной обработки ошибок в многопоточке в горутних
func safeGoRoutine(entry func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in goroutine:", r)
			// Здесь можно добавить логирование ошибки
		}
	}()
	entry()
}

func doSomething() {
	fmt.Println("something")
}

func main() {
	errChan := make(chan error)

	go func() {
		// Здесь код горутины
		if err := doSomething(); err != nil {
			errChan <- err
		}
	}()

	// Ожидание ошибки из горутины
	err := <-errChan
	if err != nil {
		// Обработка ошибки
	}
}
