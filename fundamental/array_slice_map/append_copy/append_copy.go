package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//array1 = append(array1, 4, 5, 6) // não aceita usr apend com array
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	//copy(slice2, array1) - não pode passar um array
	copy(slice2, slice1)

	fmt.Println(slice2)

}
