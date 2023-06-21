package money

import (
	"fmt"
	"strconv"
	"strings"
)

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
