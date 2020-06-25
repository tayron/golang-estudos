package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Nota: o ambiente em que esses programas são executados é determinístico, então rand.Intn sempre retornará o mesmo número.
	fmt.Println("My favorite number is", rand.Intn(10))
}
