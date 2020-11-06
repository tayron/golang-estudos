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

		var csvMatrizPessoa [][]string = [][]string{}
		listaPessoas := obterListaPessoas()

		for _, pessoa := range listaPessoas {
			pessoaArrayString := converterArrayString(pessoa)
			csvMatrizPessoa = append(csvMatrizPessoa, pessoaArrayString)
		}

		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment;filename=clientes.csv")

		wr := csv.NewWriter(w)
		for _, dado := range csvMatrizPessoa {
			err := wr.Write(dado)

			if err != nil {
				http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		wr.Flush()
	})

	fmt.Println("http://127.0.0.1:3001")
	log.Fatalln(http.ListenAndServe("127.0.0.1:3001", nil))
}

func converterArrayString(pessoa Pessoa) []string {
	var listaString []string
	listaString = append(listaString, pessoa.Nome)

	return listaString
}

func obterListaPessoas() []Pessoa {
	lista := []Pessoa{}

	lista = append(lista, Pessoa{Nome: "Pedro"})
	lista = append(lista, Pessoa{Nome: "Maria"})

	return lista
}
