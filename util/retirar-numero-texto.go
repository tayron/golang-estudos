package main

import (
	"fmt"
	"strconv"
	"strings"
)

func extractNumberFromString(str string) (int, error) {
	// Encontra o índice do primeiro dígito na string
	startIndex := strings.IndexFunc(str, func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	// Encontra o índice do último dígito na string
	endIndex := strings.LastIndexFunc(str, func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	// Verifica se há dígitos na string
	if startIndex == -1 || endIndex == -1 {
		return 0, fmt.Errorf("Nenhum número encontrado na string")
	}

	// Extrai a substring contendo o número
	numberStr := str[startIndex : endIndex+1]

	// Converte a string para um número inteiro
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("Erro ao converter número: %v", err)
	}

	return number, nil
}

func main() {
	str := `<a href="#" onclick="return OpenPageCard2(2672355);">98.04.02262698-4</a>`
	strList := strings.Split(str, ";\">")
	str = strings.Replace(strList[0], "OpenPageCard2", "OpenPageCard", 1)
	number, err := extractNumberFromString(str)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Número extraído: %d\n", number)
	}
}
