package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/tayron/go-lang-estudos/paginacao/application/handler"
)

func init() {
	godotenv.Load()
}

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	fmt.Println("Servidor startado no endereço: http://127.0.0.1:8181")
	panic(http.ListenAndServe("127.0.0.1:8181", nil))
}
