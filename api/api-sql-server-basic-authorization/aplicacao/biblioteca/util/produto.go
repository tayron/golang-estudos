package util

import (
	"errors"
	"fmt"
	"strings"
)

// ExtrairCodigoCorReferenciaProduto - Extrai refeferencia e código da cor de uma referência de um produto
func ExtrairCodigoCorReferenciaProduto(referenciaCor string) (corCodigo *string, erro error) {
	listaReferenciaCor := strings.Split(strings.Trim(referenciaCor, " "), "-")
	tamanhoListaReferenciaCor := len(listaReferenciaCor)

	if referenciaCor == "" {
		return corCodigo, errors.New("Referencia de produto vazia")
	}

	if tamanhoListaReferenciaCor == 2 || tamanhoListaReferenciaCor == 3 {
		corCodigo = &listaReferenciaCor[tamanhoListaReferenciaCor-1]
	}

	if corCodigo == nil {
		tamanhoReferencia := len(strings.Trim(referenciaCor, " "))

		if tamanhoReferencia == 9 {
			numeroCarateresRecuperar := tamanhoReferencia - 3
			var codigo string = referenciaCor[numeroCarateresRecuperar:tamanhoReferencia]
			corCodigo = &codigo
		}

		if tamanhoReferencia == 15 {
			var codigo string = referenciaCor[5:8]
			corCodigo = &codigo
		}

		if tamanhoReferencia != 9 && tamanhoReferencia != 15 {
			mensagem := fmt.Sprintf("Produto de referencia: %s não tem 9 e nem 15 caracteres, seu tamanho é de %d caracteres \n\n", referenciaCor, tamanhoReferencia)
			return corCodigo, errors.New(mensagem)
		}
	}

	return
}
