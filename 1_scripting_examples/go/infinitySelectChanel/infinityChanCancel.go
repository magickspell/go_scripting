package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("MAIN gorutine started")
	chacnelChan := make(chan struct{})
	dataChan := make(chan int)

	go func() {
		fmt.Println("second gorutine started")
		val := 0
		for {
			select {
			case <-chacnelChan:
				fmt.Println("получена отмена в канале отмены")
				return
			case dataChan <- val:
				val++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	for currentValue := range dataChan {
		fmt.Println("currentValue = ", currentValue)
		if currentValue > 42 {
			fmt.Println("отправляю отмену в канал отмены")
			chacnelChan <- struct{}{}
			break
		}
	}
}
