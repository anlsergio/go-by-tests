package api

import (
	"github.com/anlsergio/go-by-tests/webapp/model"
	"github.com/anlsergio/go-by-tests/webapp/store"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	player := "Pepper"
	db, cleanDB := createTempFile(t, "")
	defer cleanDB()

	s := &store.FileSystemPlayerStore{Database: db}
	server := NewPlayerServer(s)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatusCode(t, http.StatusOK, response.Code)
		assertResponseBody(t, "3", response.Body.String())
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatusCode(t, response.Code, http.StatusOK)

		want := []model.Player{
			{player, 3},
		}
		got := getLeagueFromResponse(t, response.Body)
		assertLeague(t, want, got)
	})
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
