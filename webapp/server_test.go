package poker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var (
	dummyBlindAlerter = &BlindAlerterSpy{}
	dummyPlayerStore  = &StubPlayerStore{}
	dummyGame         = &GameSpy{}
)

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := MustMakePlayerServer(t, &store, dummyGame)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		AssertStatus(t, http.StatusOK, spyResponse.Code)
		AssertResponseBody(t, "20", spyResponse.Body.String())
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		AssertStatus(t, http.StatusOK, spyResponse.Code)
		AssertResponseBody(t, "10", spyResponse.Body.String())
	})

	t.Run("missing player", func(t *testing.T) {
		request := newGetScoreRequest("Joseph")
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		AssertStatus(t, http.StatusNotFound, spyResponse.Code)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{scores: map[string]int{}}
	server := MustMakePlayerServer(t, store, dummyGame)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		AssertStatus(t, http.StatusAccepted, spyResponse.Code)
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
		server := MustMakePlayerServer(t, &store, dummyGame)

		request := newLeagueRequest()
		spyResponse := httptest.NewRecorder()

		server.ServeHTTP(spyResponse, request)

		got := GetLeagueFromResponse(t, spyResponse.Body)

		AssertStatus(t, http.StatusOK, spyResponse.Code)
		AssertContentType(t, spyResponse)
		AssertLeague(t, wantLeague, got)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game return 200", func(t *testing.T) {
		server := MustMakePlayerServer(t, &StubPlayerStore{}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, http.StatusOK, response.Code)
	})

	t.Run("start a game with 3 players, send blind alerts down WS, and declare Ruth the winner", func(t *testing.T) {
		wantBlindAlert := "Blind is 100"
		winner := "Ruth"

		game := &GameSpy{BlindAlert: []byte(wantBlindAlert)}
		server := httptest.NewServer(MustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, strings.Replace(server.URL, "http", "ws", 1)+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		tenMS := 10 * time.Millisecond
		time.Sleep(tenMS)

		AssertGameStartedWith(t, game, 3)
		AssertFinishCalledWith(t, game, winner)
		within(t, tenMS, func() {
			assertWebSocketGotMsg(t, ws, wantBlindAlert)
		})
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

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)

	return req
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %q %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func assertWebSocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	_, got, _ := ws.ReadMessage()

	if want != string(got) {
		t.Errorf("want %q, got %q", want, string(got))
	}
}
