package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
)

func Titulo(url string, titulo chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	re := regexp.MustCompile(`(?i)<title>\s*(.*?)\s*</title>`)
	body := make([]byte, 1024)
	resp.Body.Read(body)
	match := re.FindSubmatch(body)

	if len(match) > 1 {
		titulo <- string(match[1])
	}
}

func oMaisRapido(url1, url2, url3 string) string {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	go Titulo(url1, c1)
	go Titulo(url2, c2)
	go Titulo(url3, c3)

	// estrutura de controle específica para concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		//	default:
		//		return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)

	fmt.Println(campeao)
}
