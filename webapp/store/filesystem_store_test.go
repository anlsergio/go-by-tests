package store_test

import (
	"github.com/anlsergio/go-by-tests/webapp/model"
	"github.com/anlsergio/go-by-tests/webapp/store"
	"reflect"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	db := strings.NewReader(`[
	{"Name": "Cleo", "Wins": 10},
	{"Name": "Chris", "Wins": 33}]`)

	s := store.FileSystemPlayerStore{Database: db}

	t.Run("league from reader", func(t *testing.T) {
		want := []model.Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		got := s.GetLeague()
		assertLeague(t, want, got)

		got = s.GetLeague()
		assertLeague(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		want := 33
		got := s.GetPlayerScore("Chris")
		assertScoreEquals(t, want, got)
	})
}

func assertScoreEquals(t *testing.T, want int, got int) {
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func assertLeague(t *testing.T, want []model.Player, got []model.Player) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}
