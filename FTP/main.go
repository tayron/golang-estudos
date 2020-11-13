package main

import (
	"os"
	"time"

	"github.com/secsy/goftp"
)

func main() {

	config := goftp.Config{
		User:               "nomeusuario",
		Password:           "ICtZQAasd231",
		ConnectionsPerHost: 10,
		Timeout:            10 * time.Second,
		Logger:             os.Stderr,
	}

	client, err := goftp.DialConfig(config, "localhost")
	if err != nil {
		panic(err)
	}

	// Upload a file from disk
	bigFile, err := os.Open("caldo_cana.png")
	if err != nil {
		panic(err)
	}

	err = client.Store("caldo_cana.png", bigFile)
	if err != nil {
		panic(err)
	}
}
