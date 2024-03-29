package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/anlsergio/go-by-tests/webapp"
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

	game := poker.NewGame(store, poker.BlindAlertFunc(poker.Alerter))

	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
