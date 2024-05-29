package main

import (
	"fmt"
	"sync"
	"time"
)

// тут будет паника т.к. мапа не потоко безопасна, прикрыть ее можно mutex
// func main() {
// 	cs := map[string]int{"касса": 0}

// 	for i := 0; i < 1000; i++ {
// 		go func(k int) {
// 			cs["касса"] += 1
// 		}(i) // передаем что бы горутина работало именно с этим значением (чтобы сраблотало замыкание)
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(cs)
// }

// c mutex будет ОК
func main() {
	cs := map[string]int{"касса": 0}
	mu := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go func(k int) {
			// mu.Unlock() // тут будет паника тк поток не заглушен
			mu.Lock()
			defer mu.Unlock() // можно без дефера, но дефер выполняется в любом случае, потому лучше с ним!
			// mu.Lock() // на второй раз будет дедлок
			cs["касса"] += 1
		}(i) // передаем что бы горутина работало именно с этим значением (чтобы сраблотало замыкание)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
