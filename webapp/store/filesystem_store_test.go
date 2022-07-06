package store_test

import (
	"github.com/anlsergio/go-by-tests/webapp/model"
	"github.com/anlsergio/go-by-tests/webapp/store"
	"github.com/anlsergio/go-by-tests/webapp/tests"
	"reflect"
	"testing"
)

func TestNewFileSystemStore(t *testing.T) {
	db, cleanDB := tests.CreateTempFile(t, "")
	defer cleanDB()

	_, err := store.NewFileSystemStore(db)
	tests.AssertNoError(t, err)
}

func TestFileSystemStoreRead(t *testing.T) {
	db, cleanDB := tests.CreateTempFile(t, `[
	{"Name": "Cleo", "Wins": 10},
	{"Name": "Chris", "Wins": 33}]`)
	defer cleanDB()

	s, err := store.NewFileSystemStore(db)
	tests.AssertNoError(t, err)

	t.Run("sorted league data by score", func(t *testing.T) {
		want := []model.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		got := s.GetLeague()
		assertLeague(t, want, got)

		// read again (Seek testing)
		got = s.GetLeague()
		assertLeague(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		want := 33
		got := s.GetPlayerScore("Chris")
		assertScoreEquals(t, want, got)
	})
}

func TestFileSystemStoreWrites(t *testing.T) {
	db, cleanDB := tests.CreateTempFile(t, `[
	{"Name": "Cleo", "Wins": 10},
	{"Name": "Chris", "Wins": 33}]`)
	defer cleanDB()

	s, err := store.NewFileSystemStore(db)
	tests.AssertNoError(t, err)

	t.Run("store wins for existing players", func(t *testing.T) {
		s.RecordWin("Chris")

		want := 34
		got := s.GetPlayerScore("Chris")
		assertScoreEquals(t, want, got)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		s.RecordWin("Pepper")

		want := 1
		got := s.GetPlayerScore("Pepper")
		assertScoreEquals(t, want, got)
	})
}

func assertScoreEquals(t testing.TB, want int, got int) {
	t.Helper()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func assertLeague(t testing.TB, want []model.Player, got []model.Player) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}
