package library

import (
	"errors"
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

const estruturaContainerMenu = "<nav><ul class='pagination justify-content-center'>%s</ul></nav>Você está na página %d, exibindo %d de %d registros"
const estruturaItemMenu = "<li class='page-item'><a class='page-link' href='%s'>%s</a></li>"
const estruturaItemMenuSelecionado = "<li class='page-item active'><a class='page-link' href='%s'>%s</a></li>"
const estruturaItemMenuDesabilitado = "<li class='page-item disabled'><a class='page-link' href='%s'>%s</a></li>"

// CriarPaginacao -
func CriarPaginacao(uri string, numeroTotalRegistro int, paginaAtual int) (htmlMenu string, offset int, err error) {
	numeroRegistroPorPagina, _ := strconv.Atoi(os.Getenv("NUMERO_REGISTRO_POR_PAGINA"))

	if numeroTotalRegistro == 0 {
		return "", 0, errors.New("Numero de registro a ser paginado é zero")
	}

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

	if paginaAtual > numeroPaginasParaSeremCriadas {
		return "", 0, errors.New("Página solicitada é maior do que existente")
	}

	var numeroLinksQuePodeSerCriado int = 10 + paginaAtual
	var numeroMenuCriado int = paginaAtual
	var linksCriados string = ""

	linksCriados = fmt.Sprintf(estruturaItemMenuDesabilitado, "", "«")

	if paginaAtual > 1 {
		link = fmt.Sprintf("%s%d", paginacao.Link, 1)
		linksCriados = fmt.Sprintf(estruturaItemMenu, link, "Primeiro")
		link = fmt.Sprintf("%s%d", paginacao.Link, paginaAtual-1)
		linksCriados += fmt.Sprintf(estruturaItemMenu, link, "«")
	}

	for true {

		if numeroMenuCriado > numeroLinksQuePodeSerCriado {
			break
		}

		if numeroMenuCriado >= numeroPaginasParaSeremCriadas {
			break
		}

		link := fmt.Sprintf("%s%d", paginacao.Link, numeroMenuCriado)
		if paginaAtual == numeroMenuCriado {
			linksCriados += fmt.Sprintf(estruturaItemMenuSelecionado, link, strconv.Itoa(numeroMenuCriado))
		} else {
			linksCriados += fmt.Sprintf(estruturaItemMenu, link, strconv.Itoa(numeroMenuCriado))
		}

		numeroMenuCriado++
	}

	if paginaAtual < numeroPaginasParaSeremCriadas {
		link = fmt.Sprintf("%s%d", paginacao.Link, paginaAtual+1)
		linksCriados += fmt.Sprintf(estruturaItemMenu, link, "»")
		link = fmt.Sprintf("%s%d", paginacao.Link, numeroPaginasParaSeremCriadas)
		linksCriados += fmt.Sprintf(estruturaItemMenu, link, "Último")
	} else {
		linksCriados += fmt.Sprintf(estruturaItemMenuDesabilitado, "", "»")
	}

	offset = (paginaAtual - 1) * numeroRegistroPorPagina
	return fmt.Sprintf(estruturaContainerMenu, linksCriados, paginaAtual, numeroRegistroPorPagina, numeroTotalRegistro), offset, nil
}
