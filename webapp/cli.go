package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	input       *bufio.Scanner
}

func NewCLI(store PlayerStore, r io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		input:       bufio.NewScanner(r),
	}
}

func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.playerStore.RecordWin(extractWinner(userInput))
}

func (c *CLI) readLine() string {
	c.input.Scan()
	return c.input.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
