package main

import (
	poker "github.com/anlsergio/go-by-tests/webapp"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, closeStore, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("coult not listen on port 8080, %v", err)
	}
}
