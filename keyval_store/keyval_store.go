package keyvalstore

// Implements a Key-Value based In-Memory Database

import (
	"errors"
	"sync"
)

type Store struct {
	mu   sync.RWMutex
	Data map[string]interface{}
}

func New() *Store {
	store := &Store{
		Data: make(map[string]interface{}),
	}
	return store
}

func (s *Store) Get(key string) (interface{}, error) {

	// Prevents concurrent read/writes
	// Only Allows concurrent reads
	s.mu.Lock()
	// When function is done, unlock
	defer s.mu.Unlock()

	val, ok := s.Data[key]
	if !ok {
		err := errors.New("Error: No Key-Value pair found with key")
		return nil, err
	}

	return val, nil
}

func (s *Store) Set(key string, value interface{}) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Data[key] = value
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Data, key)
}
