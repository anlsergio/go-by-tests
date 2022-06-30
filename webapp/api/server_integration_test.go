package api

import (
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

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertStatusCode(t, http.StatusOK, response.Code)
	assertResponseBody(t, "3", response.Body.String())
}
