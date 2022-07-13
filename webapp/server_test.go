package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

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

		assertStatusCode(t, http.StatusOK, spyResponse.Code)
		assertResponseBody(t, "20", spyResponse.Body.String())
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		assertStatusCode(t, http.StatusOK, spyResponse.Code)
		assertResponseBody(t, "10", spyResponse.Body.String())
	})

	t.Run("missing player", func(t *testing.T) {
		request := newGetScoreRequest("Joseph")
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		assertStatusCode(t, http.StatusNotFound, spyResponse.Code)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{}}
	server := NewPlayerServer(&store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		assertStatusCode(t, http.StatusAccepted, spyResponse.Code)

		const wantLength = 1
		gotLength := len(store.winCalls)

		if wantLength != gotLength {
			t.Errorf("want %d calls to RecordWin, but got %d", wantLength, gotLength)
		}

		wantPlayer := player
		gotPlayer := store.winCalls[0]

		if wantPlayer != gotPlayer {
			t.Errorf("did not store the correct winner, want %q, got %q", wantPlayer, gotPlayer)
		}
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

		got := getLeagueFromResponse(t, spyResponse.Body)

		assertStatusCode(t, http.StatusOK, spyResponse.Code)
		assertContentType(t, spyResponse)
		assertLeague(t, wantLeague, got)
	})
}

func assertContentType(t testing.TB, spyResponse *httptest.ResponseRecorder) {
	if spyResponse.Result().Header.Get("content-type") != "application/json" {
		t.Error("response header does not have content-type of application/json, got ", spyResponse.Result().Header)
	}
}

func assertLeague(t *testing.T, want []Player, got []Player) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) []Player {
	t.Helper()
	league, _ := NewLeague(body)

	return league
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

func assertStatusCode(t *testing.T, want int, got int) {
	t.Helper()

	if want != got {
		t.Errorf("want status code %d, got %d", want, got)
	}
}

func assertResponseBody(t *testing.T, want string, got string) {
	t.Helper()

	if want != got {
		t.Errorf("response body is wrong, want %q, got %q", want, got)
	}
}
