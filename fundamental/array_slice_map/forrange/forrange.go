package main

import "fmt"

func main() {

	// copilador que vai contar a quantidade de elementos e defirnir no array
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numeros {

		fmt.Printf("%d) %d\n", i, numero)
	}

	// Ignorando indice do array
	for _, numero := range numeros {
		fmt.Println(numero)
	}

	// adicionando um unico atributo você pega só indice
	for indice := range numeros {
		fmt.Println(indice)
	}
}
