package main

import (
	"log"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
