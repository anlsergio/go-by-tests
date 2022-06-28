package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{
		store: &store,
	}

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

func newGetScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprint("/players/", player), nil)

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
		t.Errorf("want %q, got %q", want, got)
	}
}
