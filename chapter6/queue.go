package chapter6

import "sync"

// Queue the queue of items
type Queue struct {
	items []string
	lock  sync.RWMutex
}

// New creates a new Queue
func NewQueue() *Queue {
	q := &Queue{}
	q.items = []string{}
	return q
}

// Enqueue adds a person to the end of the queue
func (s *Queue) Enqueue(t string) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Dequeue removes a person from the start of the queue
func (s *Queue) Dequeue() *string {
	s.lock.Lock()
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	s.lock.Unlock()
	return &item
}

// Front returns the person next in the queue, without removing it
func (s *Queue) Front() *string {
	s.lock.RLock()
	int := s.items[0]
	s.lock.RUnlock()
	return &int
}

// IsEmpty returns true if the queue is empty
func (s *Queue) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the queue
func (s *Queue) Size() int {
	return len(s.items)
}
