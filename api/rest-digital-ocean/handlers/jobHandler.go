package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tayron/golang-estudos/api/rest-digital-ocean/models"
)

// GetJobs A handler to fetch all the jobs
func GetJobs(w http.ResponseWriter, r *http.Request) {
	// make a slice to hold our jobs data
	var jobs []models.Job

	// Add some job to the slice
	jobs = append(jobs, models.Job{ID: 1, Name: "Accounting"})
	jobs = append(jobs, models.Job{ID: 2, Name: "Programming"})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

// PostJob A handler to fetch all the jobs
func PostJob(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	clientData := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&clientData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var id float64 = clientData["id"].(float64)
	var name string = clientData["name"].(string)

	fmt.Println(id, name)

	switch v := clientData["age"].(type) {
	case nil:
		fmt.Println("Valor não informado")
	case int:
		fmt.Printf("Valor (%v) do tipo int informado \n", v)
	case float64:
		fmt.Printf("Valor (%v) do tipo float64 informado \n", v)
	case string:
		fmt.Printf("Valor (%v) do tipo string informado \n", v)
	default:
		fmt.Println("Tipo desconhecido")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Sucesso")

}
