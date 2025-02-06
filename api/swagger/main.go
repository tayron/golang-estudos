package main

import (
	"log"
	"net/http"

	"exemple.com/internal"
	"github.com/rs/cors"
)

func main() {

	log.Printf("Servidor startado: http://localhost:8080")

	router := internal.NewRouter()
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
