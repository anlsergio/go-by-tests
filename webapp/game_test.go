package poker_test

import (
	poker "github.com/anlsergio/go-by-tests/webapp"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	t.Run("it schedules Alerts on game start for 5 players", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.BlindAlerterSpy{}
		game := poker.NewGame(playerStore, blindAlerter)

		game.Start(5)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("it schedules Alerts on game start for 7 players", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.BlindAlerterSpy{}
		game := poker.NewGame(playerStore, blindAlerter)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewGame(store, dummyBlindAlerter)
	winner := "Ruth"

	game.Finish(winner)

	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(t *testing.T, cases []poker.ScheduledAlert, blindAlerter *poker.BlindAlerterSpy) {
	for i, want := range cases {
		t.Run(want.String(), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, want, got)
		})
	}
}

func assertScheduledAlert(t testing.TB, want, got poker.ScheduledAlert) {
	t.Helper()

	gotAmount := got.Amount
	if gotAmount != want.Amount {
		t.Errorf("want Amount %d, got %d", want.Amount, gotAmount)
	}

	gotScheduledTime := got.At
	if gotScheduledTime != want.At {
		t.Errorf("want schedule time of %v, got %v", want.At, gotScheduledTime)
	}
}
