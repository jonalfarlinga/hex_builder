package common

import "sync"

var ComponentIDS = NewSerial(0)

type Serial struct {
	mu    sync.Mutex
	value int
}

// NewSerial initializes a new Serial starting from 1 (or any custom value)
func NewSerial(start int) *Serial {
	return &Serial{
		value: start,
	}
}

// Next returns the next available ID
func (s *Serial) Next() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value++
	return s.value
}

// Current returns the current ID without incrementing
func (s *Serial) Current() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value
}
