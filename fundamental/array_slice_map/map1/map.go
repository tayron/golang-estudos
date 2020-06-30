package main

import "fmt"

func main() {
	// Mapas devem ser inicializados
	//var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[987654321] = "Pedro"
	aprovados[321321321] = "Paulo"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[987654321])

	delete(aprovados, 987654321)

	fmt.Println(aprovados[987654321])
}
