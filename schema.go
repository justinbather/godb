package godb

import (
	"errors"
	"fmt"
	"sync"
)

type Store struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewStorer() *Store {
	store := &Store{data: make(map[string][]byte)}
	return store
}

func (s *Store) ListContents() {
	for key, val := range s.data {
		fmt.Println(key, val)
	}
}

func (s *Store) Get(k string) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[k]

	if !ok {
		err := errors.New("Error: No value found")
		return nil, err
	}

	return value, nil
}

func (s *Store) Set(k string, v []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[k] = v

	return nil
}
