package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Executando...")

	log.Fatalln(http.ListenAndServe(":3001", nil))
}
