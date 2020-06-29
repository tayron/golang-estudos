package main

import "fmt"

func main() {
	i := 1

	// criando uma variavel com um endereço de memoria
	var p *int = nil

	// Pegando endereço da variavel i
	p = &i

	// Desreferenciando (pegando valor e incrementando)
	*p++
	i++

	// Go não tem operação aritimética com ponteiro+
	//p++

	fmt.Println(p, *p, i, &i)
}
