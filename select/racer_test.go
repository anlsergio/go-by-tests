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
			t.Error("error is unexpected but got :", err)
		}

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("must error out if the response takes more than the specified timeout", func(t *testing.T) {
		server := newDelayedServer(21 * time.Millisecond)
		defer server.Close()

		timeout := 20 * time.Millisecond
		_, err := _select.ConfigurableRacer(server.URL, server.URL, timeout)

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
