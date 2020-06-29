package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// Não tem while em go
	for i <= 20 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	// Laço infinito
	for {
		fmt.Println("Para sempre...")
		//time.Sleep(time.Second)
		time.Sleep(time.Second * 5) // de 5 em 5 segundos
	}
}
