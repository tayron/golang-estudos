package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// Switch true, procura o primeiro case que seja verdadeiro
	//switch true {
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}
