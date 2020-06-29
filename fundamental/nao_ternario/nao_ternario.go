package main

import "fmt"

// não existe operador ternario em GO
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"

	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
