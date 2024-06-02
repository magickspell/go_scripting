package metricInt

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("value not found")

// Cache - интерфейс кеша который нужно импллемитировать
type Cache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Metrics interface {
	TotalAmount() int64
}

type CacheWithMetrics interface {
	Cache
	Metrics
}

type impl struct {
	st    map[string]string
	mu    sync.Mutex
	total int64
}

func New() CacheWithMetrics {
	return &impl{
		st:    map[string]string{},
		mu:    sync.Mutex{},
		total: int64(0),
	}
}

func (i *impl) Get(key string) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	v, ok := i.st[key]
	if !ok {
		return "", ErrNotFound
	}

	return v, nil
}

func (i *impl) Set(key, value string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.st[key] = value
	i.total++
	return nil
}

func (i *impl) Delete(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.st, key)
	i.total--
	return nil
}

func (i *impl) TotalAmount() int64 {
	return i.total
}
