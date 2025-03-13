package main

import (
	"fmt"
	"sync"
)

// Stack - структура данных, реализующая стек
type Stack struct {
	lock sync.Mutex // используем Mutex для безопасного доступа в многопоточной среде
	data []int      // срез для хранения данных стека
}

// Push добавляет элемент на вершину стека
func (s *Stack) Push(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, item)
}

// Pop удаляет и возвращает верхний элемент стека
func (s *Stack) Pop() (int, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.data) == 0 {
		return 0, false // возвращаем false, если стек пуст
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, true
}

// Top возвращает верхний элемент стека без его удаления
func (s *Stack) Top() (int, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.data) == 0 {
		return 0, false // возвращаем false, если стек пуст
	}
	return s.data[len(s.data)-1], true
}

// как реализовать стек stack в голенге
func main() {
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if topValue, ok := s.Top(); ok {
		fmt.Println("Top value:", topValue) // Выводит: Top value: 30
	}

	if popValue, ok := s.Pop(); ok {
		fmt.Println("Popped value:", popValue) // Выводит: Popped value: 30
	}

	if topValue, ok := s.Top(); ok {
		fmt.Println("Top value now:", topValue) // Выводит: Top value now: 20
	}
}
