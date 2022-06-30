package main

import (
	"github.com/anlsergio/go-by-tests/webapp/api"
	"github.com/anlsergio/go-by-tests/webapp/store"
	"log"
	"net/http"
)

func main() {
	sv := api.NewPlayerServer(store.NewInMemoryPlayerStore())

	log.Fatal(http.ListenAndServe(":8080", sv))
}
