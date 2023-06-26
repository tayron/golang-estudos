package main

import (
	"net/http"

	"github.com/tayron/golang-estudos/testes-unitarios/api/method"
)

func main() {
	http.HandleFunc("/hello-world", method.HelloWorld)
	http.HandleFunc("/hello", method.HelloYou)
	http.HandleFunc("/upload", method.Upload)

	http.ListenAndServe(":8181", nil)
}
