package main

import (
	"fmt"
	"time"
)

// Recebe qualquer tipo
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))

	fmt.Println(tipo(func() {

	}))

	fmt.Println(tipo(time.Now()))
}
