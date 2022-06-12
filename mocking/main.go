package main

import (
	"hello/mocking/countdown"
	"os"
)

func main() {
	sleeper := &countdown.DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
