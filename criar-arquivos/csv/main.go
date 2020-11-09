package main

import (
	"encoding/csv"
	"log"
	"os"
)

var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}

func main() {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '|' // Change separator of , to |
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
