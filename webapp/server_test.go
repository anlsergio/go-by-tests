package poker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		AssertStatusCode(t, http.StatusOK, spyResponse.Code)
		AssertResponseBody(t, "20", spyResponse.Body.String())
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		AssertStatusCode(t, http.StatusOK, spyResponse.Code)
		AssertResponseBody(t, "10", spyResponse.Body.String())
	})

	t.Run("missing player", func(t *testing.T) {
		request := newGetScoreRequest("Joseph")
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		AssertStatusCode(t, http.StatusNotFound, spyResponse.Code)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{scores: map[string]int{}}
	server := NewPlayerServer(store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		AssertStatusCode(t, http.StatusAccepted, spyResponse.Code)
		AssertPlayerWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Trevor", 24},
		}

		store := StubPlayerStore{
			league: wantLeague,
		}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		got := GetLeagueFromResponse(t, spyResponse.Body)

		AssertStatusCode(t, http.StatusOK, spyResponse.Code)
		AssertContentType(t, spyResponse)
		AssertLeague(t, wantLeague, got)
	})
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)

	return req
}

func newGetScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprint("/players/", player), nil)

	return req
}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprint("/players/", player), nil)

	return req
}
