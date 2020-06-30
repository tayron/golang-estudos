package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// Slice não é um array, slice define pedaço de um array
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// Novo slice, mas aponta para o mesmo array
	s3 := a2[:2]
	fmt.Println(a2, s3)

	// Slice uam estrutura com tamanho e um ponteiro para o elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
