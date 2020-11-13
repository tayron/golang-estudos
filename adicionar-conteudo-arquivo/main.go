package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}

// AdicionarConteudoArquivoCSV - Adiciona conteúdo no arquivo csv
func AdicionarConteudoArquivoCSV(nomeArquivo string, dados [][]string) {
	nomeArquivoSCV := fmt.Sprintf("%s", nomeArquivo)
	file, err := os.OpenFile(nomeArquivoSCV, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Output(1, err.Error())
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '|' // Alterando separador para |
	defer writer.Flush()

	for _, dado := range dados {
		err := writer.Write(dado)

		if err != nil {
			log.Output(1, err.Error())
		}
	}
}

func main() {
	AdicionarConteudoArquivoCSV("teste.txt", data)
}
