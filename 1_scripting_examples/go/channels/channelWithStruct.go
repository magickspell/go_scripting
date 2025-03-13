package main

import (
	"fmt"
	"time"
)

// Определение структуры
type MyData struct {
	ID    int
	Name  string
	Value float64
}

func main() {
	// Создание канала для передачи значений типа MyData
	dataChannel := make(chan MyData)

	// Запуск горутины для отправки данных в канал
	go func() {
		for i := 0; i < 5; i++ {
			// Создание экземпляра структуры
			data := MyData{
				ID:    i,
				Name:  fmt.Sprintf("Name%d", i),
				Value: float64(i) * 1.5,
			}
			// Отправка структуры в канал
			dataChannel <- data
			time.Sleep(time.Second)
		}
		close(dataChannel)
	}()

	// Получение и вывод данных из канала
	for data := range dataChannel {
		fmt.Printf("Received: %+v\n", data)
	}
}
