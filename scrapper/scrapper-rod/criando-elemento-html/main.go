package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	headless := false
	// Configurações para iniciar o navegador
	brownserConfig := launcher.New().Headless(headless).MustLaunch()

	// Cria uma nova instância do Go Rod
	browser := rod.New().ControlURL(brownserConfig).MustConnect()

	// Abre uma nova página
	page := browser.MustPage()
	page.MustWindowMaximize()

	// Define o HTML que você deseja mockar
	htmlContent := `<html><body><h1>Hello, World!</h1></body></html>`

	page.MustEval(fmt.Sprintf(`function() { document.write("%s") }`, htmlContent))
	page.MustEval(`function() { alert(15) }`)

	// Tira uma captura de tela da página mockada
	page.MustScreenshot("mocked_page.png")

	WaitLoadBySecond(60)

	// Fecha o navegador
	browser.MustClose()
}

func WaitLoadBySecond(second time.Duration) {
	time.Sleep(second * time.Second)
}
