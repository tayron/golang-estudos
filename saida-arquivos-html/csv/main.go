package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/handlers"
)

type Pessoa struct {
	Nome string
}

// ConfigurarRotas -
func ConfigurarRotas() http.Handler {
	router := chi.NewRouter()
	router.Get("/api/v1/clientes", handlers.GetClientes)
	router.Get("/api/v1/cliente", handlers.GetClientesPorCPF)
	return router
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var listaPessoas [][]string = [][]string{}

		pessoaA := Pessoa{Nome: "Maria da silva"}
		pessoaB := Pessoa{Nome: "Pedro Augusto"}

		var pessoaAString []string
		pessoaAString = append(pessoaAString, pessoaA.Nome)
		listaPessoas = append(listaPessoas, pessoaAString)

		var pessoaBString []string
		pessoaBString = append(pessoaBString, pessoaB.Nome)
		listaPessoas = append(listaPessoas, pessoaBString)

		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment;filename=clientes.csv")

		wr := csv.NewWriter(w)
		for _, dado := range listaPessoas {
			fmt.Println(dado)

			err := wr.Write(dado)

			if err != nil {
				http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		wr.Flush()
	})

	// Executar passando a versão como parametro: version=5 go run *.go
	fmt.Println("http://127.0.0.1:3001")
	log.Fatalln(http.ListenAndServe("127.0.0.1:3001", nil))

}
