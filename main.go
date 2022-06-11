package main

import (
	"hello/depinjection"
	"os"
)

func main() {
	depinjection.Greet(os.Stdout, "Elodie")
}
