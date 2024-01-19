package godb

import (
	"errors"
	"sync"
	"time"
)

// GoDB is an in-memory key/value store database with a simple API that enables a time-to-live for a stored value aswell.
// Not meant to be fast or anything as good as Redis or something similar, just a fun project to do on the side

type item struct {
	value      interface{}
	expiration time.Time
	sliding    bool
	ttl        time.Duration
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

	// If sliding enabled we need to move the expiration up
	if val.sliding {
		val.expiration = time.Now().Add(val.ttl)
	}

	return val.value, nil
}

func (s *Store) Set(key string, value interface{}, ttl time.Duration, slide bool) {

	s.mu.Lock()
	defer s.mu.Unlock()

	expires := time.Now().Add(ttl)

	item := &item{value: value, expiration: expires, ttl: ttl, sliding: slide}

	s.Data[key] = item

	// Run a task on another thread that will remove this value once time.Now() > expiration
	go s.scheduleRemoval(key, ttl)
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Data, key)
}

func (s *Store) scheduleRemoval(key string, ttl time.Duration) {
	// Waits until current time is after the ttl duration
	<-time.After(ttl)
	s.Delete(key)
}
