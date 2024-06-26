package cache

import (
	"async-cache/cache"
	"context"
	"errors"
)

var ErrTimeOut = errors.New("time out")

type Cache struct {
	c *cache.Cache
}

func InitAsyncCache() *Cache {
	return &Cache{
		c: cache.InitCache(),
	}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	ch := make(chan string)
	go func() {
		defer close(ch)
		v, ok := c.c.Get(key)
		if ok {
			ch <- v
		}
	}()

	select {
	case <-ctx.Done():
		return "", ErrTimeOut
	case x, ok := <-ch:
		if ok {
			return x, nil
		}
		return "", errors.New("not found")
	}
}

func (c *Cache) Add(ctx context.Context, key, value string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Add(key, value)
	}()

	select {
	case <-ctx.Done(): // канал закрыт? (канал может закрыться по таймауту или вручную)
		return ErrTimeOut
	case <-ch: // появилось значение в канале?
		return nil
	}
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Del(key)
	}()

	select {
	case <-ctx.Done():
		return ErrTimeOut
	case <-ch:
		return nil
	}
}
