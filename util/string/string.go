package string

import (
	"crypto/sha256"
	"fmt"
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

// RemoverEspacosDuplos - Remove espaços duplos, deixando somente um espaço onde existe varios entre o texto
func RemoverEspacosDuplos(texto string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(texto, " ")
}

// GerarHashSha256 - Gera hash de uma string
func GerarHashSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
