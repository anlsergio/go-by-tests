package main

import (
	"hello/mocking/countdown"
	"os"
)

func main() {
	countdown.Countdown(os.Stdout)
}
