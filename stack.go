package submitter

import (
	"errors"
	"sync"
)

// Stack ...
type Stack struct {
	mutex sync.Mutex
	keys  []*Key
}

// NewStack ...
func NewStack(keys []*Key) *Stack {
	return &Stack{keys: keys}
}

// Push ...
func (s *Stack) Push(key *Key) error {
	if key == nil {
		return errors.New("nil key")
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.keys = append(s.keys, key)
	return nil
}

// Pop ...
func (s *Stack) Pop() (*Key, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.keys) == 0 {
		return nil, errors.New("empty keys")
	}
	key := s.keys[len(s.keys)-1]
	s.keys = s.keys[:len(s.keys)-1]
	return key, nil
}

// IsEmpty ...
func (s *Stack) IsEmpty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.keys) == 0
}

// Len ...
func (s *Stack) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.keys)
}
