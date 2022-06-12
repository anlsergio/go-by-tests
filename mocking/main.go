package main

import (
	"hello/mocking/countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{
		Duration:  1 * time.Second,
		SleepFunc: time.Sleep,
	}
	countdown.Countdown(os.Stdout, sleeper)
}
