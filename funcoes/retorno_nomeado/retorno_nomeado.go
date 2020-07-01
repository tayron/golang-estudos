package main

import "fmt"

//func trocar(p1, p2 int) (segundo, primeiro int) { Pode fazer assim também
func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1

	// return segundo, primeiro
	// return primeiro, segundo
	return // retorno limpo
}

func main() {
	x, y := trocar(1, 2)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
