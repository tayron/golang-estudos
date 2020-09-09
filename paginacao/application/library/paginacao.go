package library

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const estruturaContainerMenu = "<nav><ul class='pagination'>%s</ul></nav>Você está na página %d, exibindo %d de %d registros"
const estruturaItemMenu = "<li class='page-item'><a class='page-link' href='%s'>%s</a></li>"
const estruturaItemMenuSelecionado = "<li class='page-item active'><a class='page-link' href='%s'>%s</a></li>"
const estruturaItemMenuDesabilitado = "<li class='page-item disabled'><a class='page-link' href='%s'>%s</a></li>"

// CriarPaginacao -
func CriarPaginacao(numeroTotalRegistro int, r *http.Request) (htmlMenu string, offset int, err error) {
	numeroRegistroPorPagina, _ := strconv.Atoi(os.Getenv("NUMERO_REGISTRO_POR_PAGINA"))

	if numeroTotalRegistro == 0 {
		return "", 0, errors.New("Numero de registro a ser paginado é zero")
	}

	paginaAtual, _ := strconv.Atoi(fmt.Sprintf("%s", r.FormValue("pagina")))

	if paginaAtual == 0 {
		paginaAtual = 1
	}

	var numeroPaginasParaSeremCriadas int = (numeroTotalRegistro / numeroRegistroPorPagina)

	if paginaAtual > numeroPaginasParaSeremCriadas {
		return "", 0, errors.New("Página solicitada é maior do que existente")
	}

	var numeroLinksQuePodeSerCriado int = 10 + paginaAtual
	var numeroMenuCriado int = paginaAtual
	var linksCriados string = ""

	linksCriados = fmt.Sprintf(estruturaItemMenuDesabilitado, "", "«")

	if paginaAtual > 1 {
		linksCriados = fmt.Sprintf(estruturaItemMenu, montarLink(r, 1), "Primeiro")
		linksCriados += fmt.Sprintf(estruturaItemMenu, montarLink(r, paginaAtual-1), "«")
	}

	for true {

		if numeroMenuCriado > numeroLinksQuePodeSerCriado {
			break
		}

		if numeroMenuCriado >= numeroPaginasParaSeremCriadas {
			break
		}

		if paginaAtual == numeroMenuCriado {
			linksCriados += fmt.Sprintf(estruturaItemMenuSelecionado, montarLink(r, numeroMenuCriado), strconv.Itoa(numeroMenuCriado))
		} else {
			linksCriados += fmt.Sprintf(estruturaItemMenu, montarLink(r, numeroMenuCriado), strconv.Itoa(numeroMenuCriado))
		}

		numeroMenuCriado++
	}

	if paginaAtual < numeroPaginasParaSeremCriadas {
		linkProximaPagina := montarLink(r, paginaAtual+1)
		linksCriados += fmt.Sprintf(estruturaItemMenu, linkProximaPagina, "»")

		linkUltimaPagina := montarLink(r, numeroPaginasParaSeremCriadas)
		linksCriados += fmt.Sprintf(estruturaItemMenu, linkUltimaPagina, "Último")
	} else {
		linksCriados += fmt.Sprintf(estruturaItemMenuDesabilitado, "", "»")
	}

	offset = (paginaAtual - 1) * numeroRegistroPorPagina
	return fmt.Sprintf(estruturaContainerMenu, linksCriados, paginaAtual, numeroRegistroPorPagina, numeroTotalRegistro), offset, nil
}

func montarLink(r *http.Request, numeroPagina int) string {

	numeroPaginaString := strconv.Itoa(numeroPagina)

	queryString := r.URL.Query()
	queryString.Set("pagina", numeroPaginaString)
	r.URL.RawQuery = queryString.Encode()

	return r.URL.RequestURI()
}
