package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Serve arquivos estáticos da pasta "swagger-ui"
	fs := http.FileServer(http.Dir("swagger-ui"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Outras rotas
	router.HandleFunc("/api/hello", HelloWorldHandler).Methods("GET")
	router.HandleFunc("/api/health", HealthHandler).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./swagger-ui"))).Methods("GET")

	// Log das rotas registradas
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, _ := route.GetPathTemplate()

		fmt.Printf("Rota registrada: %s \n", pathTemplate)

		return nil
	})

	return router
}
