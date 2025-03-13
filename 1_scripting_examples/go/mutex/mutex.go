package main

import (
	"fmt"
	"sync"
)

// Структура, представляющая счетчик, с защитой доступа через мьютекс
type SafeCounter struct {
	mu    sync.Mutex // Встроенный мьютекс для управления доступом
	count int        // Общий ресурс, к которому нужен безопасный доступ
}

// Inc увеличивает счетчик на один в безопасной манере
func (c *SafeCounter) Inc() {
	c.mu.Lock()   // Блокировка мьютекса перед обращением к count
	c.count++     // Критическая секция, доступ к общему ресурсу
	c.mu.Unlock() // Разблокировка мьютекса после обращения к count
}

// Value возвращает текущее значение счетчика в безопасной манере
func (c *SafeCounter) Value() int {
	c.mu.Lock()         // Блокировка мьютекса перед чтением count
	defer c.mu.Unlock() // Автоматическая разблокировка мьютекса по завершении функции
	return c.count      // Возврат значения счетчика
}

func main() {
	c := SafeCounter{} // Инициализация безопасного счетчика

	// Запуск 1000 горутин для увеличения счетчика
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Final count:", c.Value())
}
