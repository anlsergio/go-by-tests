package store

import "github.com/anlsergio/go-by-tests/webapp/model"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() model.League
}
