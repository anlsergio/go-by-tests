package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	player := "Pepper"
	db, cleanDB := CreateTempFile(t, `[]`)
	defer cleanDB()

	s, err := NewFileSystemStore(db)
	AssertNoError(t, err)

	server := MustMakePlayerServer(t, s, dummyGame)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		AssertStatus(t, http.StatusOK, response.Code)
		AssertResponseBody(t, "3", response.Body.String())
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		AssertStatus(t, response.Code, http.StatusOK)

		want := []Player{
			{player, 3},
		}
		got := GetLeagueFromResponse(t, response.Body)
		AssertLeague(t, want, got)
	})
}
