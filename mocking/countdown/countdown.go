package countdown

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprint(w, finalWord)
}
