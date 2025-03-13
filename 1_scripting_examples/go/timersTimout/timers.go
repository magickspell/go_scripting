package main

import (
	"fmt"
	"sync"
	"time"
)

func longOperation() chan int {
	ch := make(chan int, 12)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
		ch <- 1337
		ch <- 3
		ch <- 4
		ch <- 5
		ch <- 6
		ch <- 7
		close(ch) // Закрываем канал после отправки всех значений
	}()
	return ch
}

func main() {
	// ОДИН ТАЙМАУТ
	// при 1 выполнится таймаут, при 3 выполнится операция
	// timer := time.NewTimer(1 * time.Second)
	timer := time.NewTimer(3 * time.Second)

	select {
	case <-timer.C:
		fmt.Println("таймаут")
	case <-time.After(time.Minute):
		// тоже самое что обычный таймер, только его нельзя остановить
		fmt.Println("time.After happened")
	case result := <-longOperation():
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("результат операции = ", result)
		fmt.Println("операция выполнена, останавливаем таймаут")
	}

	// ТИКЕР
	fmt.Println("запускаем тикер")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	step := 0
	chanTicker := longOperation()
	for tickTime := range ticker.C {
		step++
		chanValue, complete := <-chanTicker
		if !complete {
			fmt.Println("канал закрыт, выходим из цикла")
			break
		}
		fmt.Println("tickTime = ", tickTime, "   chanValue = ", chanValue)
		if step > 4 {
			fmt.Println("останавливаем тикер")
			ticker.Stop()
			break
		}
	}

	fmt.Println("запускаем бесконечный тикер")
	infinitTicker := time.Tick(time.Second)
	chanTickerInf := longOperation()
	stepInf := 0
	for tickTime := range infinitTicker {
		stepInf++
		chanValueInf, complete := <-chanTickerInf
		fmt.Println("step = ", stepInf, "tickTimeInf = ", tickTime, "chan value = ", chanValueInf)
		if !complete {
			fmt.Println("канал закрыт, выходим из бесконечного тикера")
			fmt.Println("(последнее значение 0 - это nil значение закрытого канала)")
			break
		}
		if stepInf >= 99 {
			break
		}
	}

	fmt.Println("запускаем функцию через время")
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("вывод функции запущенной через время (time.AfterFunc)")
	})
	time.Sleep(3 * time.Second)

	mu := sync.Mutex{}
	mu.Lock()
	mu.Unlock()
}
