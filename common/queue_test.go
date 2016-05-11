package common

// ref: https://github.com/hishboy/gocommons/blob/master/lang/queue.go

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthZero(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, 0, q.Len())
}

func TestPoll(t *testing.T) {
	q := NewQueue()
	expected := 20
	q.Push(expected)
	assert.Equal(t, 1, q.Len())
	val := q.Front()
	assert.Equal(t, expected, val)
	assert.Equal(t, 1, q.Len())
}

func TestPush(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 10; i++ {
		q.Push(nil)
		assert.Equal(t, i+1, q.Len())
	}
	assert.Equal(t, 10, q.Len())
	for i := 10; i > 0; i-- {
		assert.Equal(t, i, q.Len())
		val := q.Poll()
		assert.Equal(t, nil, val)
		assert.Equal(t, i-1, q.Len())
	}
}

func TestMixedPushes(t *testing.T) {
	q := NewQueue()

	q.Push(nil)
	assert.Equal(t, 1, q.Len())
	q.Push(1)
	assert.Equal(t, 2, q.Len())
	q.Push("foo")
	assert.Equal(t, 3, q.Len())
	q.Push([]int{1, 2, 3})
	assert.Equal(t, 4, q.Len())

	var val interface{}
	val = q.Poll()
	assert.Equal(t, 3, q.Len())
	assert.Equal(t, nil, val)

	val = q.Poll()
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, val, 1)

	val = q.Poll()
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, "foo", val)

	val = q.Poll()
	assert.Equal(t, 0, q.Len())
	assert.Equal(t, []int{1, 2, 3}, val)
}

func TestPullNil(t *testing.T) {
	q := NewQueue()
	val := q.Poll()
	assert.Equal(t, 0, q.Len())
	assert.Equal(t, nil, val)
}

func TestFront(t *testing.T) {
	q := NewQueue()
	expected := 10
	q.Push(expected)
	assert.Equal(t, 1, q.Len())
	val := q.Front()
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, expected, val)
}

func TestConcurrent(t *testing.T) {
	q := NewQueue()
	numGoRoutines := 50
	numPushes := 10000
	var wg sync.WaitGroup
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < numPushes; j++ {
				q.Push(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(t, numGoRoutines*numPushes, q.Len())

}
