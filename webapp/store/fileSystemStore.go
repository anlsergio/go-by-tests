package store

import (
	"github.com/anlsergio/go-by-tests/webapp/model"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadSeeker
}

func (s FileSystemPlayerStore) GetLeague() []model.Player {
	s.Database.Seek(0, 0)
	league, _ := model.NewLeague(s.Database)

	return league
}
