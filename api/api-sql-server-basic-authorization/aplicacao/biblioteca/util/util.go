package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ObterCaminhoDiretorioAplicacao -
func ObterCaminhoDiretorioAplicacao() string {
	ambienteAplicacao := os.Getenv("AMBIENTE")

	if ambienteAplicacao == "desenvolvimento" {
		return ObterCaminhoDiretorioAplicacaoWeb()
	}

	return ObterCaminhoDiretorioAplicacaoLinux()
}

// ObterCaminhoDiretorioAplicacaoWeb -
func ObterCaminhoDiretorioAplicacaoWeb() string {
	caminho, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return caminho
}

// ObterCaminhoDiretorioAplicacaoLinux -
func ObterCaminhoDiretorioAplicacaoLinux() string {
	caminho, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	return caminho
}

// DebugarStruct -
func DebugarStruct(dado interface{}) {
	empJSON, err := json.MarshalIndent(dado, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s\n", string(empJSON))
}
