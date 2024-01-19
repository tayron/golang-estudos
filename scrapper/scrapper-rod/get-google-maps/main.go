package main

import (
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

const (
	URL      = "https://www.google.com.br/search?sca_esv=599783327&tbs=lf:1,lf_ui:10&tbm=lcl&q=venda+de+produtos+para+manicure+no+nordeste&rflfq=1&num=10&sa=X&ved=2ahUKEwigu8yUwOmDAxUsHDQIHcZUBPAQjGp6BAgeEAE&biw=1920&bih=936&dpr=1#rlfi=hd:;si:;mv:[[-3.455797,-34.655904],[-8.3232248,-38.8183556]];tbs:lrf:!1m4!1u3!2m2!3m1!1e1!1m4!1u2!2m2!2m1!1e1!2m1!1e2!2m1!1e3!3sIAE,lf:1,lf_ui:10"
	HEADLESS = true
	TIMEOUT  = 10
)

func main() {
	browser, page := openPage(URL, HEADLESS, TIMEOUT)
	defer browser.MustClose()

	elementDiv, err := page.ElementX(`//*[@id="rl_ist0"]/div/div[1]/div[3]`)
	if err != nil {
		log.Printf("Erro ao acessar div com todas as empresas: %s", err.Error())
	}

	Listdetails, err := elementDiv.Elements(".rllt__details")
	if err != nil {
		log.Printf("Erro ao acessar div com dados das empresas: %s", err.Error())
	}

	for _, div := range Listdetails[1:] {
		if strings.TrimSpace(div.MustText()) != "" {
			title, err := div.Element("span")
			if err != nil {
				log.Printf("Erro ao obter o título: %s", err.Error())
			}

			log.Println(title.MustText())
		}
	}
}

func openPage(url string, headless bool, timeout int) (*rod.Browser, *rod.Page) {
	browser := rod.New().ControlURL(
		launcher.New().Headless(headless).MustLaunch(),
	).Timeout(time.Second * time.Duration(timeout)).MustConnect()

	page := browser.MustPage(url)

	return browser, page
}
