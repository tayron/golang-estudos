package httprequest

import (
	"encoding/json"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}
