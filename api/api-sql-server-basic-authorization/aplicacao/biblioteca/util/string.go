package util

import (
	"log"
	"regexp"
	"strings"
)

// RemoverCaracteresEspeciais - Remove todos caracteres especiais mantendo somente o que for letra e numero
func RemoverCaracteresEspeciais(texto string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(texto, "")
}

// ConverteStringUperCase - Converte nomes para padrão UpperCase, remove espaços em branco antes de depois da string e remove aspas simples
func ConverteStringUperCase(nome string) string {
	return strings.TrimSpace(strings.ToUpper(strings.ReplaceAll(nome, "'", "")))
}
