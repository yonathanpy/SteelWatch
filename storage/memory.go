package storage

import (
	"sync"
	"steelwatch/types"
)

type Store struct {
	mu     sync.Mutex
	events []types.Event
}

func NewStore() *Store {
	return &Store{
		events: make([]types.Event, 0),
	}
}

func (s *Store) Add(e types.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, e)
}

func (s *Store) All() []types.Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.events
}
