package store

import (
	"encoding/json"
	"github.com/anlsergio/go-by-tests/webapp/model"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) (wins int) {
	player := s.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}

	return wins
}

func (s *FileSystemPlayerStore) GetLeague() model.League {
	s.resetReaderOffset()
	league, _ := model.NewLeague(s.Database)

	return league
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	league := s.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, model.Player{
			Name: name,
			Wins: 1,
		})
	}

	s.resetReaderOffset()
	json.NewEncoder(s.Database).Encode(league)
}

func (s *FileSystemPlayerStore) resetReaderOffset() {
	s.Database.Seek(0, 0)
}
