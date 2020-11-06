package util

import (
	"log"
	"os"
)

// GravarLog - Utiliário para armazenamento de log em arquivos no sistema
func GravarLog(nomeArquivo, mensagem string) {
	f, err := os.OpenFile("aplicacao.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Output(1, err.Error())
	}
	defer f.Close()

	logger := log.New(f, nomeArquivo+" ", log.LstdFlags)
	logger.Print(mensagem)
}
