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

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{At: duration, Amount: amount})
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

func AssertStatusCode(t *testing.T, want int, got int) {
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
