package library

import (
	"fmt"
	"os"
	"strconv"
)

type Paginacao struct {
	NumeroTotalRegistro     int    // 60
	NumeroRegistroPorPagina int    // 20
	PaginaAtual             int    //1
	URIBase                 string // produtos/listar
	QueryString             string // pagina=
	Link                    string //produtos/listar/pagina=
}

const estruturaContainerMenu = "<nav><ul class='pagination'>%s</ul></nav>"
const estruturaItemMenu = "<li class='page-item'><a class='page-link' href='%s'>%d</a></li>"

// CriarPaginacao -
func CriarPaginacao(uri string, numeroTotalRegistro int, paginaAtual int) (htmlMenu string, offset int) {
	numeroRegistroPorPagina, _ := strconv.Atoi(os.Getenv("NUMERO_REGISTRO_POR_PAGINA"))

	var link string = "/?pagina="

	if uri != "/" {
		link = uri + "/" + "?pagina="
	}

	if paginaAtual == 0 {
		paginaAtual = 1
	}

	var paginacao Paginacao = Paginacao{
		NumeroRegistroPorPagina: numeroRegistroPorPagina,
		NumeroTotalRegistro:     numeroTotalRegistro,
		PaginaAtual:             paginaAtual,
		URIBase:                 uri,
		QueryString:             "?pagina=",
		Link:                    link,
	}

	var numeroPaginasParaSeremCriadas int = (paginacao.NumeroTotalRegistro / paginacao.NumeroRegistroPorPagina)
	var numeroMenuCriado int = 0
	var linksCriados string = ""

	for true {
		numeroMenuCriado++

		if numeroMenuCriado >= numeroPaginasParaSeremCriadas {
			break
		}

		link := fmt.Sprintf("%s%d", paginacao.Link, numeroMenuCriado)
		linksCriados += fmt.Sprintf(estruturaItemMenu, link, numeroMenuCriado)
	}

	offset = (paginaAtual - 1) * numeroRegistroPorPagina
	return fmt.Sprintf(estruturaContainerMenu, linksCriados), offset
}
