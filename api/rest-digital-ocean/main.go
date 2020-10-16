package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tayron/go-lang-estudos/api/rest-digital-ocean/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Get("/api/jobs", handlers.GetJobs)
	router.Post("/api/job", handlers.PostJob)

	// run it on port 8080
	fmt.Println("Server executing at port 8180")
	err := http.ListenAndServe("0.0.0.0:8180", router)

	if err != nil {
		log.Fatal(err)
	}
}
