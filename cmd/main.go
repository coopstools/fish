package main

import (
	"os"
	"time"

	"github.com/coopstools/fish/internal"
)

func main() {
	app := os.Args[1]
	layout := internal.Open(app)
	layout.InitPrint()
	for {
		layout.Print()
		time.Sleep(500 * time.Millisecond)
		layout.Update()
	}
}
