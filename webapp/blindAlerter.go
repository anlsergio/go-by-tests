package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlertFunc implements BlindAlerter so that users have
// the option to implement the BlindAlerter interface with a function
// rather than an empty struct.
type BlindAlertFunc func(duration time.Duration, amount int)

func (b BlindAlertFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	b(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
