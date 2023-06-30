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
	browser, page, err := OpenPage("", false, 50)
	defer browser.MustClose()

	if err != nil {
		t.Error(err.Error())
	}

	htmlContent := `<html><body><h1 id=\"title\">Hello, World!</h1></body></html>`
	page.MustEval(fmt.Sprintf(`function() { document.write("%s") }`, htmlContent))
	page.MustScreenshot("mocked_page.png")

	const MENSAGEM_ERRO = "O valor esperado %s, mas o resultado encontrado foi %s."
	valorEsperado := "<h1 id=\"title\">Hello, World!</h1>"
	valorEncontrado, errGetElement := page.Element("#title")

	if errGetElement != nil {
		t.Error(errGetElement.Error())
	}

	if valorEncontrado.MustHTML() != valorEsperado {
		t.Errorf(MENSAGEM_ERRO, valorEsperado, valorEncontrado.MustHTML())
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
		fmt.Printf("Não foi possível abrir a página inicial. err[%v] \n ", err)
		return browser, page, err
	}

	page.MustWindowMaximize()

	return browser, page, err
}

func WaitLoadBySecond(second time.Duration) {
	time.Sleep(second * time.Second)
}
