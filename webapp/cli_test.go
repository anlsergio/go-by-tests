package poker_test

import (
	"bytes"
	"fmt"
	poker "github.com/anlsergio/go-by-tests/webapp"
	"strings"
	"testing"
	"time"
)

var (
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user in", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(playerStore, dummyBlindAlerter)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user in", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(playerStore, dummyBlindAlerter)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it prompts the user for the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(dummyPlayerStore, blindAlerter)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		want := poker.NumberOfPlayersPrompt
		got := stdout.String()

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}
			})

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, want, got)
		}
	})
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
