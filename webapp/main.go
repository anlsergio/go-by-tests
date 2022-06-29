package main

import (
	"github.com/anlsergio/go-by-tests/webapp/api"
	"github.com/anlsergio/go-by-tests/webapp/store"
	"log"
	"net/http"
)

func main() {
	sv := &api.PlayerServer{
		Store: store.NewInMemoryPlayerStore(),
	}

	log.Fatal(http.ListenAndServe(":8080", sv))
}
