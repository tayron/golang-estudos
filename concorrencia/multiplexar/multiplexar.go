package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		// Sempre que chegar dados no canal de origem armazena no canal de destino
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
// juntando varias mensagens de canais diferente e juntar no mesmo canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		titulo("https://www.cod3r.com.br", "https://www.google.com"),
		titulo("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println("Primeiros:", <-c, "|", <-c)
	fmt.Println("Segundos:", <-c, "|", <-c)
}

func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			title := r.FindStringSubmatch(string(html))

			if len(title) > 0 {
				c <- r.FindStringSubmatch(string(html))[1]
			} else {
				mensagem := "Não foi possível encontrar titulo para a url: " + url
				c <- mensagem
			}
		}(url)

	}

	return c
}
