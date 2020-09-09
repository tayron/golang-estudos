package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/tayron/go-lang-estudos/paginacao/application/library"
	"github.com/tayron/go-lang-estudos/paginacao/application/view"

	"github.com/tayron/go-lang-estudos/paginacao/application/models"
)

// HomeHandler -
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	uri := "/"
	numeroTotalRegistro := models.ObterNumeroProdutos()
	queryStringPagina := r.FormValue("pagina")
	paginaAtual, _ := strconv.Atoi(fmt.Sprintf("%s", queryStringPagina))

	htmlPaginacao, offset := library.CriarPaginacao(uri, numeroTotalRegistro, paginaAtual)
	listaProdutos := models.BuscarTodosProdutos(offset)

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
