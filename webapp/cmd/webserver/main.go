package main

import (
	"log"
	"net/http"

	poker "github.com/anlsergio/go-by-tests/webapp"
)

const dbFileName = "game.db.json"

func main() {
	store, closeStore, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()

	game := poker.NewGame(store, poker.BlindAlertFunc(poker.Alerter))

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("could not create a Player Server, %v", err)
	}

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("coult not listen on port 8080, %v", err)
	}
}
