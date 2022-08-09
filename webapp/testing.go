package poker

import (
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s *ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips At %v", s.Amount, s.At)
}

type BlindAlerterSpy struct {
	Alerts []ScheduledAlert
}

func (b *BlindAlerterSpy) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	b.Alerts = append(b.Alerts, ScheduledAlert{At: duration, Amount: amount})
}

func CreateTempFile(t testing.TB, initialData string) (fileBuffer *os.File, removeTempFile func()) {
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

type GameSpy struct {
	StartCalled      bool
	StartedWith      int
	BlindAlert       []byte
	FinishCalled     string
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers

	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalledWith = winner
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}

func AssertScoreEquals(t testing.TB, want int, got int) {
	t.Helper()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func GetLeagueFromResponse(t testing.TB, body io.Reader) []Player {
	t.Helper()
	league, _ := NewLeague(body)

	return league
}

func AssertStatus(t *testing.T, want int, got int) {
	t.Helper()

	if want != got {
		t.Errorf("want status code %d, got %d", want, got)
	}
}

func AssertResponseBody(t *testing.T, want string, got string) {
	t.Helper()

	if want != got {
		t.Errorf("response body is wrong, want %q, got %q", want, got)
	}
}

func AssertContentType(t testing.TB, spyResponse *httptest.ResponseRecorder) {
	t.Helper()

	if spyResponse.Result().Header.Get("content-type") != "application/json" {
		t.Error("response header does not have content-type of application/json, got ", spyResponse.Result().Header)
	}
}

func AssertLeague(t testing.TB, want []Player, got []Player) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}

func MustMakePlayerServer(t *testing.T, store PlayerStore, game Game) *PlayerServer {
	server, err := NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func AssertFinishCalledWith(t testing.TB, game *GameSpy, wantWinner string) {
	t.Helper()

	ok := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishCalledWith == wantWinner
	})

	if !ok {
		t.Errorf("expected finish called with %q but got %q", wantWinner, game.FinishCalledWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)

	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}

	return false
}

func AssertGameStartedWith(t testing.TB, game *GameSpy, wantPlayers int) {
	t.Helper()

	gotPlayers := game.StartedWith

	if wantPlayers != gotPlayers {
		t.Errorf("expected Start called with %d players, but got %d", wantPlayers, gotPlayers)
	}
}
