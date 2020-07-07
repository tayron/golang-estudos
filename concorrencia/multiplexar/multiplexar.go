package main

import (
	"fmt"

	"github.com/tayron/go-lang-estudos/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		// Sempre que chegar dados no canal de origem armazena no canal de destino
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
// juntando varias mensagens de canais diferente e juntar no mesmo canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println("Primeiros:", <-c, "|", <-c)
	fmt.Println("Segundos:", <-c, "|", <-c)
}
