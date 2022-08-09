package poker

import (
	"fmt"
	"io"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlertFunc implements BlindAlerter so that users have
// the option to implement the BlindAlerter interface with a function
// rather than an empty struct.
type BlindAlertFunc func(duration time.Duration, amount int, to io.Writer)

func (b BlindAlertFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	b(duration, amount, to)
}

func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
