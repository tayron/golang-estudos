package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4                   // Buffer lotado -> make(chan int, 3) buffer de tamanho 3
	fmt.Println("Executou!!") // Não é executado porque fica esperando liberar espaço no buff e armazenar o valor 4
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // e ler o 3º libera espaço no buffer para o ch <- 4
	fmt.Println(<-ch)
}
