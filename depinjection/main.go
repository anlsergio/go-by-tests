package main

import (
	"hello/depinjection/greet"
	"log"
	"net/http"
	"os"
)

func main() {
	greet.Greet(os.Stdout, "Elodie")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(greet.GreetHandler)))
}
