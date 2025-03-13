package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик горутин в группе

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second * 3) // Имитация длительной работы
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)         // Увеличиваем счетчик горутин в группе
		go worker(i, &wg) // Запуск горутины
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("All workers done")
}
