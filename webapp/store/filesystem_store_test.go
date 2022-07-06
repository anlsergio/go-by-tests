package store_test

import (
	"github.com/anlsergio/go-by-tests/webapp/model"
	"github.com/anlsergio/go-by-tests/webapp/store"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemStoreRead(t *testing.T) {
	db, cleanDB := createTempFile(t, `[
	{"Name": "Cleo", "Wins": 10},
	{"Name": "Chris", "Wins": 33}]`)
	defer cleanDB()

	s := store.FileSystemPlayerStore{Database: db}

	t.Run("league from reader", func(t *testing.T) {
		want := []model.Player{
			{"Cleo", 10},
			{"Chris", 33},
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
	db, cleanDB := createTempFile(t, `[
	{"Name": "Cleo", "Wins": 10},
	{"Name": "Chris", "Wins": 33}]`)
	defer cleanDB()

	s := store.FileSystemPlayerStore{Database: db}

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

func createTempFile(t testing.TB, initialData string) (fileBuffer io.ReadWriteSeeker, removeTempFile func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempFile.Write([]byte(initialData))

	removeTempFile = func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return tempFile, removeTempFile
}
