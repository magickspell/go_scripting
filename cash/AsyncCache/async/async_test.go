package cache

import (
	"context"
	"testing"
	"time"
)

func TestAsyncSet(t *testing.T) {
	ac := InitAsyncCache()
	to := time.Millisecond

	ctxBase := context.Background()            // сделали пустой контекст (инициализировали)
	ctx, _ := context.WithTimeout(ctxBase, to) // контекст с таймаутом 1 милисекунда
	err := ac.Add(ctx, "k", "v")
	if err != ErrTimeOut {
		t.Error("expected timeout")
	}

	to = time.Millisecond * 2
	ctx, _ = context.WithTimeout(ctxBase, to)
	err = ac.Add(ctx, "k", "v")
	if err != nil {
		t.Errorf("expected Set. %v", err)
	}
}

func TestAsyncGet(t *testing.T) {
	ac := InitAsyncCache()
	to := time.Millisecond
	key := "k"
	value1 := "v1"

	ctxBase := context.Background()            // сделали пустой контекст (инициализировали)
	ctx, _ := context.WithTimeout(ctxBase, to) // контекст с таймаутом 1 милисекунда

	_ = ac.Add(ctx, key, value1)
	_, err := ac.Get(ctx, key)
	if err != ErrTimeOut {
		t.Error("Expected timeout!")
	}

	ctx, _ = context.WithTimeout(ctxBase, to*5)
	v, err := ac.Get(ctx, key)
	if err != nil {
		t.Error("expected Get.")
	}
	if v != value1 {
		t.Errorf("expected <%v>, got <%v>", value1, v)
	}
}
