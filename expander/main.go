package main

import (
	"fmt"
	"html/template"
	"math/big"
	"net/http"

	"github.com/tayron/golang-estudos/expander/expanders"
)

// Produto -
type Produto struct {
	Nome  string
	Preco *big.Float
}

func createTemplate(fileName string) *template.Template {
	pathToFile := fmt.Sprintf("./pages/%s.html", fileName)
	templateBase := template.New("base.html").Funcs(expanders.All())
	tplHome, error := templateBase.ParseFiles("./pages/base.html", pathToFile)
	return template.Must(tplHome, error)
}

// HomeHandler -
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	parameters := struct {
		PageTitle string
		Produtos  []Produto
	}{
		PageTitle: "Produtos",
		Produtos: []Produto{
			{Nome: "Arroz", Preco: big.NewFloat(15.23)},
			{Nome: "Feijão", Preco: big.NewFloat(5.21)},
		},
	}

	template := createTemplate("home")
	template.Execute(w, parameters)
}

func main() {

	http.HandleFunc("/", HomeHandler)
	fmt.Println("Servidor startado no endereço: http://127.0.0.1:8182")
	panic(http.ListenAndServe("127.0.0.1:8182", nil))

}
