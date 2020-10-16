package main

import (
	"fmt"
	"unsafe"
)

type carro struct {
	nome            string
	velocidadeAtual int
}

// composição
type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Printf("A ferrari %s está com turbo ligado? %v \n", f.nome, f.turboLigado)
	fmt.Println(f)

	// extrai só a parte referente ao struct carro
	bar := *(*carro)(unsafe.Pointer(&f))
	fmt.Printf("%+v\n", bar)
}
