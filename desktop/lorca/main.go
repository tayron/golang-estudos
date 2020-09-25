package main

import (
	"log"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("http://www.google.com.br", "", 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	ui.SetBounds(lorca.Bounds{
		WindowState: "fullscreen",
	})
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
