package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // Enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base // Enviando dados para o canal

	time.Sleep(3 * time.Second)
	c <- 4 * base // Enviando dados para o canal
}

func main() {

	// Channel (canal) é a forma de comunicação entre grroutines
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	fmt.Println("A")

	a, b := <-c, <-c // Recebendo dados do canal
	fmt.Println("B")

	fmt.Println(a, b)
	fmt.Println(<-c)

}
