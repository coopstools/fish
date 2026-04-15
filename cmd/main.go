package main

import (
	"time"

	"github.com/coopstools/fish/internal"
)

func main() {
	layout := internal.Open("./example/0_around.fish")
	layout.InitPrint()
	for {
		layout.Print()
		time.Sleep(500 * time.Millisecond)
		layout.Update()
	}
}
