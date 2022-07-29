package poker_test

import (
	"bytes"
	poker "github.com/anlsergio/go-by-tests/webapp"
	"strings"
	"testing"
)

var (
	dummyBlindAlerter = &poker.BlindAlerterSpy{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user in", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("5\nChris wins\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		wantPrompt := poker.NumberOfPlayersPrompt
		gotPrompt := stdout.String()

		if wantPrompt != gotPrompt {
			t.Errorf("want %q, got %q", wantPrompt, gotPrompt)
		}

		wantPlayers := 5
		gotPlayers := game.StartedWith

		if wantPlayers != gotPlayers {
			t.Errorf("expected Start called with %d players, but got %d", wantPlayers, gotPlayers)
		}

		wantWinner := "Chris"
		gotWinner := game.FinishedWith

		if wantWinner != gotWinner {
			t.Errorf("want %q winner, but got %q", wantWinner, gotWinner)
		}
	})

	t.Run("it prompts the user for the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		wantPrompt := poker.NumberOfPlayersPrompt
		gotPrompt := stdout.String()

		if wantPrompt != gotPrompt {
			t.Errorf("want %q, got %q", wantPrompt, gotPrompt)
		}

		wantPlayers := 7
		gotPlayers := game.StartedWith

		if wantPlayers != gotPlayers {
			t.Errorf("expected Start called with %d players, but got %d", wantPlayers, gotPlayers)
		}
	})
}
