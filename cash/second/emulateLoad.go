package main

import (
  "errors"
  "fmt"
  "sync"
)

func emulateLoad(c Cache, parallelFactor int) {
  wg := sync.WaitGroup{}

  for i := 0; i < parallelFactor; i++ {
    // с этим ключ/значением будем работать на этой итерации цикла
    key := fmt.Sprintf("#{i}-key")
    value := fmt.Sprintf("#{i}-value")

    wg.Add(1)
    // запсиь в кеш
    go func(k string) {
      err := c.Set(k, value)
      if err != nil {
        panic(err)
      }
      wg.Done()
    }(key)

    wg.Add(1)
    // чтение из кеша
    go func(k, v string) {
      _, err := c.Get(k)
      // Если другая горутина не успела удалить значение из кеша
      // Проверим что оно совпадает с тем что мы хотелиб добавить в кеш
      if err != nil && !errors.Is(err, ErrNotFound) {
        panic(err)
      }
      wg.Done()
    }(key, value)

    wg.Add(1)
    // Удаление из кеша
    go func(k string) {
      err := c.Delete(k)
      if err != nil {
        panic(err)
      }
      wg.Done()
    }(key)

    // ждем пока отработают все горутины
    wg.Wait()
  }
}

func emulateReadIntensiveLoad(c Cache, parallelFactor int) {
  wg := sync.WaitGroup{}

  for i := 0; i < parallelFactor/10; i++ {
    // с этим ключ/значением будем работать на этой итерации цикла
    key := fmt.Sprintf("#{i}-key")
    value := fmt.Sprintf("#{i}-value")

    wg.Add(1)
    // запсиь в кеш
    go func(k string) {
      err := c.Set(k, value)
      if err != nil {
        panic(err)
      }
      wg.Done()
    }(key)

    wg.Add(1)
    // Удаление из кеша
    go func(k string) {
      err := c.Delete(k)
      if err != nil {
        panic(err)
      }
      wg.Done()
    }(key)

    // ждем пока отработают все горутины
    wg.Wait()
  }

  for i := 0; i < parallelFactor; i++ {
    // с этим ключ/значением будем работать на этой итерации цикла
    key := fmt.Sprintf("#{i}-key")
    value := fmt.Sprintf("#{i}-value")

    wg.Add(1)
    // чтение из кеша
    go func(k, v string) {
      _, err := c.Get(k)
      // Если другая горутина не успела удалить значение из кеша
      // Проверим что оно совпадает с тем что мы хотелиб добавить в кеш
      if err != nil && !errors.Is(err, ErrNotFound) {
        panic(err)
      }
      wg.Done()
    }(key, value)
  }
}
