package main

import "errors"

func main() {

}

var ErrNotFound = errors.New("value not found")

// Cache - интерфейс кеша который нужно импллемитировать
type Cache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

// 1connectionSimple реализация
/*
Мапа НЕ потокобезопасная, мутекса нет - плохо.
данные не будут синхронизированы.
*/
type simpleCache struct {
	storage map[string]string
}

func NewCache() Cache {
	return &simpleCache{
		storage: make(map[string]string),
	}
}

func (c *simpleCache) Set(key string, value string) error {
	c.storage[key] = value

	return nil
}

func (c *simpleCache) Get(key string) (string, error) {
	value, ok := c.storage[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (c *simpleCache) Delete(key string) error {
	delete(c.storage, key)

	return nil
}
