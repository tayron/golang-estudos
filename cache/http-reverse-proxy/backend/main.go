package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var counter uint64

func productsHandler(w http.ResponseWriter, r *http.Request) {
	hit := atomic.AddUint64(&counter, 1)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=30")

	payload := map[string]any{
		"product":      "notebook-gamer",
		"hit_number":   hit,
		"generated_at": time.Now().Format(time.RFC3339Nano),
	}

	fmt.Println("backend gerou nova resposta, hit:", hit)
	json.NewEncoder(w).Encode(payload)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", productsHandler)

	fmt.Println("backend ouvindo na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
