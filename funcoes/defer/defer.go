package main

import "fmt"

func obtertValorAprovado(numero int) int {
	defer fmt.Println("fim!")

	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}

	fmt.Println("valor baixoo...")
	return numero

}

func main() {
	fmt.Println(obtertValorAprovado(6000))
	fmt.Println(obtertValorAprovado(3000))
}
