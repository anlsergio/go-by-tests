package store

import (
	"github.com/anlsergio/go-by-tests/webapp/model"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadSeeker
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range s.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}

func (s *FileSystemPlayerStore) GetLeague() []model.Player {
	s.Database.Seek(0, 0)
	league, _ := model.NewLeague(s.Database)

	return league
}
