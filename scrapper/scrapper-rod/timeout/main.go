package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	const TIMEOUT = 0
	const PAGINA = "https://www.hospeda.app"

	var html string

	err := rod.Try(func() {
		html = getPageHTML(PAGINA, TIMEOUT)
	})

	if errors.Is(err, context.DeadlineExceeded) {
		html = fmt.Sprintf("Página %s demorou mais que %d segundos para abrir \n", PAGINA, TIMEOUT)
	} else if err != nil {
		html = fmt.Sprintf("Erro desconhecido: %s \n", err.Error())
	}

	fmt.Println(html)
}

func getPageHTML(pagina string, timeout time.Duration) string {
	page := rod.New().MustConnect().MustPage()
	page.Timeout(timeout * time.Second).MustNavigate(pagina)
	return page.MustHTML()
}
