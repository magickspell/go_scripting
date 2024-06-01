package main

import (
	"sync"
	"sync/atomic"
	"time"
)

func MutexCounter() int {
	gorutinesCount := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			gorutinesCount++
			m.Unlock()
			time.Sleep(time.Microsecond)
			m.Lock()
			gorutinesCount--
			m.Unlock()
			wg.Done()
		}()
	}
	// мы ожидаем что счетчик останется равным нулю
	wg.Wait()
	return gorutinesCount
}

func AtomicCounter() int32 {
	gorutinesCount := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&gorutinesCount, 1)
			time.Sleep(time.Microsecond)
			atomic.AddInt32(&gorutinesCount, -1)
			wg.Done()
		}()
	}
	// мы ожидаем что счетчик останется равным нулю
	wg.Wait()
	return gorutinesCount
}
