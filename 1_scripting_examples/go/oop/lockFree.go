package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int64
	atomic.AddInt64(&counter, 1) // Атомарно увеличиваем счетчик на 1
	fmt.Println("Counter:", counter)
}
