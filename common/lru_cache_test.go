package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Set("1", 1)
	cache.Set("2", 2)
	cache.Set("3", 3)
	assert.Equal(t, 2, cache.Len())
	assert.Equal(t, 2, cache.MaxSize())

	_, ok := cache.Get("2")
	assert.Equal(t, true, ok)
	_, ok = cache.Get("3")
	assert.Equal(t, true, ok)
	_, ok = cache.Get("1")
	assert.Equal(t, false, ok)

	cache.Delete("2")
	_, ok = cache.Get("2")
	assert.Equal(t, false, ok)
	assert.Equal(t, 1, cache.Len())
}
