package keyvalstore

// Implements a Key-Value based In-Memory Database

import (
	"errors"
	//"fmt"
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expiration time.Time
}

type Store struct {
	mu   sync.RWMutex
	Data map[string]*item
}

func New() *Store {
	store := &Store{
		Data: make(map[string]*item),
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

	return val.value, nil
}

func (s *Store) Set(key string, value interface{}, ttl time.Duration) {

	s.mu.Lock()
	defer s.mu.Unlock()

	expires := time.Now().Add(ttl)

	item := &item{value: value, expiration: expires}

	s.Data[key] = item

	go s.scheduleRemoval(key, ttl)
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Data, key)
}

func (s *Store) scheduleRemoval(key string, ttl time.Duration) {
	<-time.After(ttl)
	s.Delete(key)
}
