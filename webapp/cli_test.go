package poker_test

import (
	"bytes"
	poker "github.com/anlsergio/go-by-tests/webapp"
	"io"
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
	t.Run("start game with 3 players and finish with 'Chris' as a winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}

		in := userInput("3", "Chris wins")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.NumberOfPlayersPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishWithWinner(t, game, "Chris")
	})

	t.Run("start game with 8 players and finish with 'Cleo' as a winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}

		in := userInput("8", "Cleo wins")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.NumberOfPlayersPrompt)
		assertGameStartedWith(t, game, 8)
		assertFinishWithWinner(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered", func(t *testing.T) {
		game := &poker.GameSpy{}

		in := userInput("Pies")
		stdout := &bytes.Buffer{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.NumberOfPlayersPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when an invalid winner command is entered", func(t *testing.T) {
		game := &poker.GameSpy{}

		in := userInput("3", "Lloyd is a killer")
		stdout := &bytes.Buffer{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.NumberOfPlayersPrompt, poker.BadWinnerInputErrMsg)
	})
}

func userInput(in ...string) io.Reader {
	return strings.NewReader(strings.Join(in, "\n"))
}

func assertGameNotStarted(t testing.TB, game *poker.GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("the game should not have started due to invalid input")
	}
}

func assertFinishWithWinner(t *testing.T, game *poker.GameSpy, wantWinner string) {
	gotWinner := game.FinishedWith

	if wantWinner != gotWinner {
		t.Errorf("want %q winner, but got %q", wantWinner, gotWinner)
	}
}

func assertGameStartedWith(t testing.TB, game *poker.GameSpy, wantPlayers int) {
	t.Helper()

	gotPlayers := game.StartedWith

	if wantPlayers != gotPlayers {
		t.Errorf("expected Start called with %d players, but got %d", wantPlayers, gotPlayers)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, msgs ...string) {
	t.Helper()

	want := strings.Join(msgs, "")
	got := stdout.String()

	if want != got {
		t.Errorf("want %q sent to stdout, but got %q", want, got)
	}
}
