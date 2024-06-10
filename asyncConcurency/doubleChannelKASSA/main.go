package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"касса #1connectionSimple": 0, "касса #2": 0}
	mu := sync.Mutex{}

	for i := 0; i < 500; i++ {
		go func(k int) {
			mu.Lock()
			defer mu.Unlock()
			cs["касса #1connectionSimple"] += 1
		}(i)
	}

	for i := 0; i < 1000; i++ {
		go func(k int) {
			mu.Lock()
			defer mu.Unlock() // можно без дефера, но дефер выполняется в любом случае, потому лучше с ним!
			cs["касса #2"] += 1
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
