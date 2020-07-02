package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, nome := range aprovados {
		fmt.Printf("%d) %s \n", i+1, nome)
	}
}

func main() {

	aprovados := []string{"Maria", "Pedro", "Paulo"}
	imprimirAprovados(aprovados...)

}
