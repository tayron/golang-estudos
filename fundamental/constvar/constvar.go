package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float(64) inferido pelo complicador

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circuferência é", area)

	const (
		a = 2
		b = 3
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
