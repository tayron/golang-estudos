package main

import "log"

func init() {
	//SetTimezone("America/Sao_Paulo")
	SetTimezone("America/Los_Angeles")
}

func main() {
	for true {
		log.Print("Teste")
	}
}
