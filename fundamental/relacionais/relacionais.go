package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings", "bananas" == "bananas")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas", d1 == d2)

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Pedro"}
	p2 := Pessoa{"Pedro"}

	fmt.Println("Pessoas", p1 == p2)
}
