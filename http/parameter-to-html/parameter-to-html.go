package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)

	appVersion := fmt.Sprintf(os.Getenv("version"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := struct {
			AppVersion string
		}{
			AppVersion: appVersion,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Executar passando a versão como parametro: version=5 go run *.go
	fmt.Println("http://127.0.0.1:3001")
	log.Fatalln(http.ListenAndServe("127.0.0.1:3001", nil))

}
