package main

import "fmt"

// TODO: Refatorar esta função para evitar código duplicado
func add(a, b int) int {
	return a + b
}

// Função não utilizada para gerar alerta
func unusedFunction() {
	fmt.Println("This function is not used")
}

func main() {
	fmt.Println("Hello, SonarQube!")
	result := add(2, 3)
	fmt.Println("Result:", result)

	// Código duplicado para gerar alerta
	sum := add(4, 5)
	fmt.Println("Sum:", sum)
}
