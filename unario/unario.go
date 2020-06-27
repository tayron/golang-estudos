package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas pos fixada
	x++
	fmt.Println(x)

	y--
	fmt.Println(y)

	// não é permitido fazer acrescimo ou descrecimo dentro de comparação
	//fmt.Println(x == y--)

}
