package store_test

import (
	"context"
	"errors"
	store "hello/context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	t        *testing.T
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	datac := make(chan string, 1)

	go s.SlowlyFetchResponse(ctx, datac)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case responseData := <-datac:
		return responseData, nil
	}
}

func (s *SpyStore) SlowlyFetchResponse(ctx context.Context, datac chan string) {
	var result string
	for _, c := range s.response {
		select {
		case <-ctx.Done():
			log.Println("spy server got cancelled")
		default:
			time.Sleep(5 * time.Millisecond)
			result += string(c)
		}
	}
	datac <- result
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	want := "hello, world"
	s := &SpyStore{
		t:        t,
		response: want,
	}
	svr := store.Server(s)

	t.Run("returns store server response", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != want {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), want)
		}
	})

	t.Run("tells store server to cancel work if request is cancelled", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(ctx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("data should not be written to the response at this point")
		}
	})
}
