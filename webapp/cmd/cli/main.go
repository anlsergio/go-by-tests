package main

import (
	"fmt"
	poker "github.com/anlsergio/go-by-tests/webapp"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker!")
	fmt.Println("Type '{Name} wins' to record a win")

	store, closeStore, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
