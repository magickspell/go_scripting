package main

import (
	"fmt"
)

// Простая хеш-функция
func simpleHash(key int) int {
	return key % 5 // остаток от деления на 5
}

// Элемент хеш-таблицы
type Entry struct {
	Key   int
	Value string
}

// Хеш-таблица с использованием связного списка для разрешения коллизий
type HashTable struct {
	buckets [5][]Entry // 5 бакетов для упрощения примера
}

// Добавление элемента в хеш-таблицу
func (h *HashTable) Put(key int, value string) {
	index := simpleHash(key)
	// Проверяем, существует ли уже ключ в бакете
	for i, entry := range h.buckets[index] {
		if entry.Key == key {
			h.buckets[index][i].Value = value // Обновление существующего значения
			return
		}
	}
	// Добавление новой записи, если ключ уникален
	h.buckets[index] = append(h.buckets[index], Entry{Key: key, Value: value})
}

// Получение значения по ключу
func (h *HashTable) Get(key int) (string, bool) {
	index := simpleHash(key)
	for _, entry := range h.buckets[index] {
		if entry.Key == key {
			return entry.Value, true
		}
	}
	return "", false // Ключ не найден
}

func main() {
	hashTable := &HashTable{}
	hashTable.Put(1, "Значение 1")
	fmt.Println("hashTable значение:", hashTable)
	hashTable.Put(6, "Значение 6")
	fmt.Println("hashTable значение:", hashTable)
	hashTable.Put(11, "Значение 11")
	hashTable.Put(12, "Значение 11") // кроме 12

	fmt.Println("hashTable значение:", hashTable)
	// все запишетсчя в бакет под ключом 1, кроме 12

	// Попытка получить значения по ключам
	if value, found := hashTable.Get(6); found {
		fmt.Println("Найдено значение:", value)
	} else {
		fmt.Println("Значение не найдено.")
	}
}
