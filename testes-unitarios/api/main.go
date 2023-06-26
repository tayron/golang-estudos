package main

import (
	"net/http"

	"github.com/tayron/golang-estudos/testes-unitarios/api/method"
)

func main() {
	http.HandlerFunc("/hello-world", method.HelloWorld)
	http.HandlerFunc("/hello", method.HelloYou)
	http.HandlerFunc("/upload", method.Upload)

	http.ListenAndServe(":80", nil)
}
