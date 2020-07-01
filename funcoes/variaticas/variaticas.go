package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("%.2f \n", media())
	fmt.Printf("%.2f \n", media(9, 8, 5))
	fmt.Printf("%.2f \n", media(9, 5))
	fmt.Printf("%.2f \n", media(9, 5, 9, 2, 4, 6))
}
