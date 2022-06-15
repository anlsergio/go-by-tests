package counter_test

import (
	counter "hello/sync"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3x leaves it at 3", func(t *testing.T) {
		c := counter.NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()

		want := 3

		assertCounter(t, want, c)
	})

	t.Run("concurrent safe for race conditions", func(t *testing.T) {
		want := 1000
		c := counter.NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, want, c)
	})
}

func assertCounter(t *testing.T, want int, c *counter.Counter) {
	t.Helper()

	got := c.Value()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
