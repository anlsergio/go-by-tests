package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game *Game
}

const NumberOfPlayersPrompt = "Please enter the number of players: "

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, NumberOfPlayersPrompt)

	numberOfPlayersInput := c.readLine()
	numberOfPlayers, _ := strconv.Atoi(numberOfPlayersInput)

	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	winner := extractWinner(winnerInput)

	c.game.Finish(winner)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func NewCLI(in io.Reader, out io.Writer, game *Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
