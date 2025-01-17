package money

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseMoneyIntToString(value int64) string {
	// Verifica se o valor é negativo
	isNegative := value < 0
	if isNegative {
		value = -value // Remove o sinal temporariamente para formatar
	}

	// Divide por 100 para obter o valor em reais com duas casas decimais
	reais := float64(value) / 100.0
	formatted := fmt.Sprintf("%.2f", reais) // Ex: "123456.78"

	// Substitui o ponto decimal por vírgula
	formatted = strings.Replace(formatted, ".", ",", 1)

	// Separa a parte inteira e decimal
	parts := strings.Split(formatted, ",")
	intPart := parts[0]
	decimalPart := parts[1]

	// Adiciona pontos na parte inteira
	var withDots strings.Builder
	for i, digit := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			withDots.WriteRune('.')
		}
		withDots.WriteRune(digit)
	}

	// Adiciona o prefixo "R$" e o sinal negativo, se necessário
	result := fmt.Sprintf("R$ %s,%s", withDots.String(), decimalPart)
	if isNegative {
		result = fmt.Sprintf("R$ -%s", result[3:]) // Adiciona o sinal negativo antes do número
	}

	return result
}

func ParseMoneyStringToFloat(value string) (moneyFormated float64, err error) {
	value = strings.TrimPrefix(value, "R")
	value = strings.TrimPrefix(value, "$")
	value = strings.TrimSpace(value)

	numDot := strings.Count(value, ".")
	numComma := strings.Count(value, ",")

	if numDot == 0 && numComma == 1 {
		value = strings.ReplaceAll(value, ",", ".")
		return parseFloat(value)
	}

	if numDot == 1 && numComma == 0 || numDot == 0 && numComma == 0 {
		return parseFloat(value)
	}

	if numDot >= 1 && numComma == 1 {
		value = strings.ReplaceAll(value, ".", "")
		value = strings.ReplaceAll(value, ",", ".")
		return parseFloat(value)
	}

	if numDot == 2 && numComma == 0 {
		firstDotIndex := strings.Index(value, ".")
		if firstDotIndex != -1 {
			value = value[:firstDotIndex] + value[firstDotIndex+1:]
		}

		return parseFloat(value)
	}

	return
}

func parseFloat(value string) (float64, error) {
	valueFormated, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("Erro ao converter string '%s' para float", value)
	}

	return valueFormated, nil
}
