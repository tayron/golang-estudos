package main

import "fmt"

func worker(workerID int, canal chan int) {
	for x := range canal {
		fmt.Printf("Worker %d received %d\n", workerID, x)
	}
}

func main() {
	canal := make(chan int)
	quantidadeWorkers := 100

	// Inicializa os workers
	for i := 0; i < quantidadeWorkers; i++ {
		go worker(i, canal)
	}

	// Joga carga para or workers
	for i := 0; i < 1000000; i++ {
		canal <- i
	}
}
