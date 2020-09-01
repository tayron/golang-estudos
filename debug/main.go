package main

import "fmt"

func main() {
	var nome string
	var sobrenome string
	var nomeCompleto string

	nome = "Tayron"
	sobrenome = "Miranda"
	nomeCompleto = fmt.Sprintf("%s %s", nome, sobrenome)

	fmt.Println(nomeCompleto)

}
