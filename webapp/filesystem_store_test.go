package poker

import (
	"testing"
)

func TestNewFileSystemStore(t *testing.T) {
	db, cleanDB := CreateTempFile(t, "")
	defer cleanDB()

	_, err := NewFileSystemStore(db)
	AssertNoError(t, err)
}

func TestFileSystemStoreRead(t *testing.T) {
	db, cleanDB := CreateTempFile(t, `[
	{"Name": "Cleo", "Wins": 10},
	{"Name": "Chris", "Wins": 33}]`)
	defer cleanDB()

	s, err := NewFileSystemStore(db)
	AssertNoError(t, err)

	t.Run("sorted league data by score", func(t *testing.T) {
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		got := s.GetLeague()
		AssertLeague(t, want, got)

		// read again (Seek testing)
		got = s.GetLeague()
		AssertLeague(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		want := 33
		got := s.GetPlayerScore("Chris")
		AssertScoreEquals(t, want, got)
	})
}

func TestFileSystemStoreWrites(t *testing.T) {
	db, cleanDB := CreateTempFile(t, `[
	{"Name": "Cleo", "Wins": 10},
	{"Name": "Chris", "Wins": 33}]`)
	defer cleanDB()

	s, err := NewFileSystemStore(db)
	AssertNoError(t, err)

	t.Run("playerStore wins for existing players", func(t *testing.T) {
		s.RecordWin("Chris")

		want := 34
		got := s.GetPlayerScore("Chris")
		AssertScoreEquals(t, want, got)
	})

	t.Run("playerStore wins for new players", func(t *testing.T) {
		s.RecordWin("Pepper")

		want := 1
		got := s.GetPlayerScore("Pepper")
		AssertScoreEquals(t, want, got)
	})
}
