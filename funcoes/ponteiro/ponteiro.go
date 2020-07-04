package main

import "fmt"

func inc1(n int) {
	n++
}

// um ponteiro é representado por um *
// recebe um ponteiro como parametro
func inc2(n *int) {
	// * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	//& é uasdo para obter o endereço de um variavel
	inc2(&n) // por referencia
	fmt.Println(n)
}
