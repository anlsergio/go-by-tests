package main

import (
	poker "github.com/anlsergio/go-by-tests/webapp"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("could not open file %s, %v", dbFileName, err)
	}

	st, err := poker.NewFileSystemStore(db)
	if err != nil {
		log.Fatalf("problem creating player store, %v", err)
	}

	sv := poker.NewPlayerServer(st)

	if err := http.ListenAndServe(":8080", sv); err != nil {
		log.Fatalf("coult not listen on port 8080, %v", err)
	}
}
