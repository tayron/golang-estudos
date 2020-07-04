package main

import "fmt"

// Este método é executado sozinho pelo go antes de tudo
func init() {
	fmt.Println("Inicializando..")
}

// go run *.go
func main() {
	fmt.Println("Main..n")
}
