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

	timeout := 0
	i := 0

	for {
		html := ""
		i++
		err := rod.Try(func() {
			html = getPageHTML(PAGINA, timeout)
		})

		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("Página %s demorou mais que %d segundos para abrir \n", PAGINA, timeout)
			timeout = 1
			continue
		} else if err != nil {
			fmt.Printf("Erro desconhecido: %s \n", err.Error())
			timeout = 1
			continue
		}

		if html != "" {
			fmt.Printf("OK: %d \n", i)
		}

		timeout = 0
	}

}

func getPageHTML(pagina string, timeout int) string {
	timeoutSeconds := time.Second * time.Duration(timeout)
	page := rod.New().MustConnect().MustPage()
	page.Timeout(timeoutSeconds).MustNavigate(pagina)
	return page.MustHTML()
}
