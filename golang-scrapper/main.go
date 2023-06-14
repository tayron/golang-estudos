package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	fmt.Println("Testando scrapper com GO-Rod")
	page := rod.New().MustConnect().MustPage()

	page.MustNavigate("https://www.hospeda.app")

	html := page.MustHTML()
	fmt.Println("HTML DA PAGINA:", html)
}
