package main

import (
	"hello/maths/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
