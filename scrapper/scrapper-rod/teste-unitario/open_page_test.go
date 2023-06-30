package testeunitario

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func TestOpenPage(t *testing.T) {
	browser, page, err := OpenPage("", true, 50)
	defer browser.MustClose()

	if err != nil {
		t.Error(err.Error())
	}

	htmlContent := `<html><body><h1 id=\"title\">Hello, World!</h1></body></html>`
	page.MustEval(fmt.Sprintf(`function() { document.write("%s") }`, htmlContent))

	valorEsperado := "<h1 id=\"title\">Hello, World!</h1>"
	valorEncontrado, errGetElement := page.Element("#title")

	if errGetElement != nil {
		t.Error(errGetElement.Error())
	}

	if valorEncontrado.MustHTML() != valorEsperado {
		mensagem := "O valor esperado %s, mas o resultado encontrado foi %s."
		t.Errorf(mensagem, valorEsperado, valorEncontrado.MustHTML())
	}
}

func TestOpenPageWithTimeout(t *testing.T) {
	urlSite := "http://www.oi.com.br/"
	timeoutAtSeconds := 0

	errTimeout := rod.Try(func() {
		browser, page, errorOpenPage := OpenPage(urlSite, true, timeoutAtSeconds)
		defer browser.MustClose()

		if errorOpenPage == nil {
			t.Errorf("Site deveria ter dado timeout em %d segundos", timeoutAtSeconds)
		}

		pageTitle := page.MustElement("title").MustHTML()
		if pageTitle != "" {
			t.Errorf("O site não deveria ter sido aberto em %d segundos", timeoutAtSeconds)
		}
	})

	if errTimeout == nil {
		t.Error("Timout não aconteceu")
	}
}

func TestOpenPageWithoutTimeout(t *testing.T) {
	urlSite := "http://oi.com.br"
	timeoutAtSeconds := 15

	errTimeout := rod.Try(func() {
		browser, page, errorOpenPage := OpenPage(urlSite, true, timeoutAtSeconds)
		defer browser.MustClose()

		if errorOpenPage != nil && errorOpenPage.Error() == fmt.Sprintf("não foi possível abrir a pagina: %s", urlSite) {
			t.Error(errorOpenPage.Error())
		}

		pageTitle := page.MustElement("title").MustHTML()
		if pageTitle == "" {
			t.Errorf("Não foi possível abrir o site do linkedin em %d segundos", timeoutAtSeconds)
		}
	})

	if errTimeout != nil {
		t.Error("Aconteceu timeout, sendo que não deveria ter ocorrido")
	}
}

func OpenPage(url string, headless bool, timeout int) (*rod.Browser, *rod.Page, error) {
	brownserConfig := launcher.New().Headless(headless).MustLaunch()
	browser := rod.New()
	browser = browser.ControlURL(brownserConfig)
	browser = browser.Timeout(time.Second * time.Duration(timeout))
	browser = browser.MustConnect()

	opt := proto.TargetCreateTarget{
		URL: url,
	}

	page, err := browser.Page(opt)
	if err != nil {
		return browser, page, fmt.Errorf("não foi possível abrir a pagina: %s", url)
	}

	page.MustWindowMaximize()

	return browser, page, err
}
