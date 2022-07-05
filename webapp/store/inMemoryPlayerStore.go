package store

import "github.com/anlsergio/go-by-tests/webapp/model"

type InMemoryPlayerStore struct {
	store map[string]int
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.store[name]++
}

func (s *InMemoryPlayerStore) GetLeague() (league model.League) {
	for name, wins := range s.store {
		league = append(league, model.Player{
			Name: name,
			Wins: wins,
		})
	}

	return
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
	}
}
