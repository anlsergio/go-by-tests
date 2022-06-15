package store

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(s Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.Fetch(r.Context())
		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
