package poker

import (
	"github.com/anlsergio/go-by-tests/webapp/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	player := "Pepper"
	db, cleanDB := tests.CreateTempFile(t, `[]`)
	defer cleanDB()

	s, err := NewFileSystemStore(db)
	tests.AssertNoError(t, err)

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

		want := []Player{
			{player, 3},
		}
		got := getLeagueFromResponse(t, response.Body)
		assertLeague(t, want, got)
	})
}
