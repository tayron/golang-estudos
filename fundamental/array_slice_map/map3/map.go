package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel":  1500.00,
			"Gabriela": 2300.00,
		},
		"P": {
			"Paulo": 1530.21,
			"Pedro": 1360.00,
		},
		"C": {
			"Carla": 1530.21,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	for _, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
