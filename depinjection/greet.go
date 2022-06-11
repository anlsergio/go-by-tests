package depinjection

import (
	"fmt"
	"io"
	"os"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
