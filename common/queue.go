package common

import (
	"sync"
)

type queuenode struct {
	data interface{}
	next *queuenode
}

// Queue is a FIFO data structure
type Queue struct {
	head  *queuenode
	tail  *queuenode
	count int
	lock  *sync.Mutex
}

// NewQueue creates Queue
func NewQueue() *Queue {
	q := &Queue{}
	q.lock = &sync.Mutex{}
	return q
}

// Len returns length of Queue
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.count
}

// Push pushes item to the Queue
func (q *Queue) Push(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := &queuenode{data: item}

	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
}

// Poll returns and removes the top item in Queue
func (q *Queue) Poll() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next

	if q.head == nil {
		q.tail = nil
	}
	q.count--

	return n.data
}

// Front returns the top item in Queue
// This operation does NOT modify queue
func (q *Queue) Front() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := q.head
	if q.head == nil {
		return nil
	}

	return n.data
}
