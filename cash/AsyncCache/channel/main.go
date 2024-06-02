package channel

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	cc := Init()
	wg.Add(1)

	for i := 0; i < 10; i++ {
		ch := cc.GetChannel("a")
		fmt.Printf("Gorutings %v\n", runtime.NumGoroutine())
		fmt.Printf("ch len %v\n", len(ch))
	}
	time.Sleep(time.Second)
	fmt.Printf("after sleep Gorutings %v\n", runtime.NumGoroutine())
}
