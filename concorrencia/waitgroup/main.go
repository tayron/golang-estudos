package main

import (
	"fmt"
	"sync"
)

func process(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Indica que a goroutine terminou

	// Simulação de processamento
	fmt.Printf("Goroutine %d: Iniciada\n", i)
	for j := 0; j < 100000000; j++ {
		// Processamento pesado
	}
	fmt.Printf("Goroutine %d: Concluída\n", i)
}

func main() {
	var wg sync.WaitGroup

	// Número de goroutines que queremos executar
	numGoroutines := 5

	// Adicionar o número de goroutines à WaitGroup
	wg.Add(numGoroutines)

	// Iniciar as goroutines
	for i := 0; i < numGoroutines; i++ {
		go process(i, &wg)
	}

	// Esperar até que todas as goroutines terminem
	wg.Wait()

	fmt.Println("Todas as goroutines foram concluídas.")
}
