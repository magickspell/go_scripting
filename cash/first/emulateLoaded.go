package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func emulateLoaded(t *testing.T, c Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		// с этим ключ/значением будем работать на этой итерации цикла
		key := fmt.Sprintf("#{i}-key")
		value := fmt.Sprintf("#{i}-value")

		wg.Add(1)
		// запсиь в кеш
		go func(k string) {
			err := c.Set(k, value)
			assert.NoError(t, err) // если условие не выполнено - тест стопнется
			wg.Done()
		}(key)

		wg.Add(1)
		// чтение из кеша
		go func(k, v string) {
			storedValue, err := c.Get(k)
			// Если другая горутина не успела удалить значение из кеша
			// Проверим что оно совпадает с тем что мы хотелиб добавить в кеш
			if !errors.Is(err, ErrNotFound) {
				assert.Equal(t, v, storedValue)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		// Удаление из кеша
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			wg.Done()
		}(key)

		// ждем пока отработают все горутины
		wg.Wait()
	}
}
