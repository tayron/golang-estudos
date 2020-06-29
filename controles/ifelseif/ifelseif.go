package main

import "fmt"

func nota(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota <= 9 {
		return "B"
	} else if nota >= 3 && nota <= 5 {
		return "C"
	} else {
		return "D"
	}
}

func main() {
	fmt.Println(nota(9.7))
	fmt.Println(nota(8.5))
	fmt.Println(nota(4.0))
	fmt.Println(nota(2.3))
}
