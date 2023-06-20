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

	page := rod.New().MustConnect().MustPage()
	var html string

	err := rod.Try(func() {
		page.Timeout(TIMEOUT * time.Second).MustNavigate(PAGINA)
		html = page.MustHTML()
	})

	if errors.Is(err, context.DeadlineExceeded) {
		html = fmt.Sprintf("Página %s demorou mais que %d segundos para abrir \n", PAGINA, TIMEOUT)
	} else if err != nil {
		html = fmt.Sprintf("Erro desconhecido: %s \n", err.Error())
	}

	fmt.Println(html)
}
