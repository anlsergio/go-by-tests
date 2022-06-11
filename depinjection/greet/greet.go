package greet

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
