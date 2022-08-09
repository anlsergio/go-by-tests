package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	NumberOfPlayersPrompt = "Please enter the number of players: "
	BadPlayerInputErrMsg  = "you are so silly!"
	BadWinnerInputErrMsg  = "wrong call pal"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, NumberOfPlayersPrompt)

	numberOfPlayersInput := c.readLine()

	numberOfPlayers, err := strconv.Atoi(numberOfPlayersInput)
	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}

	c.game.Start(numberOfPlayers, c.out)

	winnerInput := c.readLine()
	winner := extractWinner(winnerInput)
	if winner == "" {
		fmt.Fprint(c.out, BadWinnerInputErrMsg)
		return
	}

	c.game.Finish(winner)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func extractWinner(userInput string) string {
	winner := strings.Replace(userInput, " wins", "", 1)

	if len(winner) == len(userInput) {
		return ""
	}

	return winner
}
