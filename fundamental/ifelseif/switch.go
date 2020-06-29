package main

import "fmt"

func nota(nota float64) string {
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota <= 9:
		return "B"
	case nota >= 3 && nota <= 5:
		return "C"
	default:
		return "D"
	}
}

func main() {
	fmt.Println(nota(9.7))
	fmt.Println(nota(8.5))
	fmt.Println(nota(4.0))
	fmt.Println(nota(2.3))
}
