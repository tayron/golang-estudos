package main

import (
	"fmt"

	"github.com/tayron/golang-estudos/testes-unitarios/mock/mock-interface/entity"
)

func main() {
	calculadora := entity.Calculadora{}
	calculadora.Somar(10, 5)

	fmt.Println(calculadora.Resultado)

	calculadoraMock = entity.CalculadoraMock{}
	calculadora.Somar(10, 5)
	fmt.Println(calculadora.Resultado)
}
