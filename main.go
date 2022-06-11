package main

import (
	"hello/depinjection"
	"log"
	"net/http"
	"os"
)

func main() {
	depinjection.Greet(os.Stdout, "Elodie")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(depinjection.GreetHandler)))
}
