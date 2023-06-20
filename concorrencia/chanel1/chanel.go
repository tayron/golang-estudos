package main

import "fmt"

func main() {
	ch := make(chan int, 1) // criando um canal de inteiro no tamanho 1

	ch <- 1 // Enviando dados para o canal (escrita)
	// Recebendo dados do canal (leitura)
	fmt.Println(<-ch)

	ch <- 2
	fmt.Println(<-ch)

}
