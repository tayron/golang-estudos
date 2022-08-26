package handler

import (
	"html/template"
	"net/http"

	"github.com/tayron/golang-estudos/paginacao/application/library"
	"github.com/tayron/golang-estudos/paginacao/application/view"

	"github.com/tayron/golang-estudos/paginacao/application/models"
)

// HomeHandler -
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	numeroTotalRegistro := models.ObterNumeroProdutos()
	htmlPaginacao, offset, err := library.CriarPaginacao(numeroTotalRegistro, r)

	var listaProdutos []models.Produto

	if err == nil {
		listaProdutos = models.BuscarTodosProdutos(offset)
	}

	parameters := struct {
		PageTitle string
		Produtos  []models.Produto
		Paginacao template.HTML
	}{
		PageTitle: "Produtos",
		Produtos:  listaProdutos,
		Paginacao: template.HTML(htmlPaginacao),
	}

	view.LoadViewTemplate("home", w, parameters)
}
