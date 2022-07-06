package store

import (
	"encoding/json"
	"github.com/anlsergio/go-by-tests/webapp/model"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	Database io.Writer
	League   model.League
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) (wins int) {
	player := s.League.Find(name)
	if player != nil {
		return player.Wins
	}

	return wins
}

func (s *FileSystemPlayerStore) GetLeague() model.League {
	return s.League
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	player := s.League.Find(name)

	if player != nil {
		player.Wins++
	} else {
		s.League = append(s.League, model.Player{
			Name: name,
			Wins: 1,
		})
	}

	json.NewEncoder(s.Database).Encode(s.League)
}

func NewFileSystemStore(db *os.File) *FileSystemPlayerStore {
	db.Seek(0, 0)
	league, _ := model.NewLeague(db)

	return &FileSystemPlayerStore{
		Database: &tape{db},
		League:   league,
	}
}
