package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	mu := &sync.Mutex{}
	mu.Lock()
	mu.Unlock()

	value := 1
	value++
	fmt.Println(value)

	value2 := int64(1)
	atomic.AddInt64(&value2, 1) // атомик гарантированно прибавит 1 (сделает операцию), но нужна строгая типизация
	fmt.Println(value2)

	// ctx := context.Background()
	// value3 := atomic.Value{}
	// value3.Store(&ctx)
	// av := value3.Load().(context.Context)
	// fmt.Println(av)

	value4 := atomic.Value{}
	value4.Store(100)
	av4 := value4.Load()
	fmt.Println(av4)
}
