package main

import (
	"fmt"

	lib "github.com/tayron/golang-estudos/workscape/calculadora-lib"
)

func main() {
	var a = 5
	var b = 15
	var resultado = lib.Soma(a, b)
	fmt.Printf("Resultado da soma de %d + %d é %d \n", a, b, resultado)
}
