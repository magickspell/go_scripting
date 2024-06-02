package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cache(t *testing.T) {
	t.Parallel()
	testCache := NewCache()

	t.Run("correctly stored value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)

		assert.Equal(t, value, storedValue)
	})

	t.Run("correctly updated value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedValue)

		newValue := "newValue"
		err = testCache.Set(key, newValue)
		newStoredValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, newValue, newStoredValue)
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()
		parallelFactor := 100_000
		emulateLoad(testCache, parallelFactor)
	})
}
