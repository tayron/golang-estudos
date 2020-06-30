package main

import "fmt"

func main() {
	funcsSalarios := map[string]float64{
		"José":  1300.35,
		"Pedro": 3500.00,
		"Maria": 9512.25,
	}

	funcsSalarios["Paulo"] = 6521.23

	delete(funcsSalarios, "adfad")

	for nome, salario := range funcsSalarios {
		fmt.Println(nome, salario)
	}
}
