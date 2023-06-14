package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	fmt.Println("Testando scrapper com GO-Rod")

	browser, page := openPage("https://www.hospeda.app", true, 10)
	defer browser.MustClose()

	html := page.MustHTML()
	fmt.Println("HTML DA PAGINA:", html)
}

func openPage(url string, headless bool, timeout int) (*rod.Browser, *rod.Page) {
	browser := rod.New().ControlURL(
		launcher.New().Headless(true).MustLaunch(),
	).Timeout(time.Second * time.Duration(timeout)).MustConnect()

	page := browser.MustPage(url)

	return browser, page
}
