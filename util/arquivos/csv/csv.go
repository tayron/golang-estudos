package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// CriarArquivoCSV - Cria e gravar arquivo .csv com determinado conteúdo
func CriarArquivoCSV(nomeArquivo string, dados [][]string) {
	nomeArquivoSCV := fmt.Sprintf("%s/%s.csv", os.Getenv("DIRETORIO_CRIACAO_ARQUIVOS_CSV"), nomeArquivo)
	file, err := os.Create(nomeArquivoSCV)

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
