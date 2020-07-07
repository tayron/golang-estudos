package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 180) // 180 milissegundos
				break
			}
		}
	}

	close(c) // Sempre feche um canal para não gerar o deadLock
}

func main() {
	c := make(chan int, 30) // canal de inteiro com buffer de 30 posições

	//go primos(60, c)
	go primos(cap(c), c) // cap de capacidade, tamanho buff

	for primo := range c {
		fmt.Printf("%d ", primo)
	}

	fmt.Println("Fim!")
}
