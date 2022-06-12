package _select_test

import (
	_select "hello/select"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("quickest response wins", func(t *testing.T) {
		slowServer := newDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := newDelayedServer(0)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := _select.Racer(slowURL, fastURL)
		if err != nil {
			t.Error("unexpected error but got :", err)
		}

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("must timeout if the response takes more than 10s", func(t *testing.T) {
		serverA := newDelayedServer(11 * time.Second)
		defer serverA.Close()
		serverB := newDelayedServer(12 * time.Second)
		defer serverB.Close()

		_, err := _select.Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("error expected but didn't get any")
		}
	})
}

func newDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
