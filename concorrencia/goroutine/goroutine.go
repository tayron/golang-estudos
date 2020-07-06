package main

import (
	"fmt"
	"time"
)

func fale(pessoa, text string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d) \n", pessoa, text, i+1)
	}
}

func main() {

	// Funções que sao executadas de forma independêntes

	//fale("Maria", "Porque você não fala comigo?", 3)
	//fale("João", "Só posso falar depois de você!", 1)

	//go fale("Maria", "Ei ...", 500)
	//go fale("João", "Opa...", 500)

	//time.Sleep(time.Second * 5)

	//fmt.Println("Fim!")

	go fale("Maria", "Etendi..", 10)
	fale("João", "Parabéns!!", 5)
}
