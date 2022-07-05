package api

import (
	"github.com/anlsergio/go-by-tests/webapp/model"
	"github.com/anlsergio/go-by-tests/webapp/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	player := "Pepper"
	server := NewPlayerServer(store.NewInMemoryPlayerStore())

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