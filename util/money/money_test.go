package money

import "testing"

func TestParseMoneyStringToFloatSuccess(t *testing.T) {
	t.Parallel()

	const MENSAGEM_ERRO = "Valor esperado é %.2f, porém o valor retornado foi %.2f"

	listaValores := map[string]float64{
		"R$ 9.111.356.224,01": 9111356224.01,
		"$ 956.451.966,91":    956451966.91,
		"R$ 1.356.224,01":     1356224.01,
		"R$ 1.224,01":         1224.01,
		"156.320.19":          156320.19,
		"1.200.00":            1200.00,
		"1.200.01":            1200.01,
		"1.235.12":            1235.12,
		"15.21":               15.21,
		"37":                  37,
		"0.13":                0.13,
		"1,23":                1.23,
	}

	for valorParaTestar, valorEsperado := range listaValores {
		valorConvertido, err := ParseMoneyStringToFloat(valorParaTestar)

		if err != nil {
			t.Errorf("%s", err.Error())
		}

		if valorConvertido != valorEsperado {
			t.Errorf(MENSAGEM_ERRO, valorEsperado, valorConvertido)
		}
	}
}

func TestParseMoneyStringToFloatWithWrongValue(t *testing.T) {
	t.Parallel()

	const MENSAGEM_ERRO = "O valor esperado '%s', mas o resultado encontrado foi '%s'."

	valor := "casa"
	mensagemErro := "Erro ao converter string 'casa' para float"
	_, erro := ParseMoneyStringToFloat(valor)

	if erro.Error() != mensagemErro {
		t.Errorf(MENSAGEM_ERRO, mensagemErro, erro.Error())
	}
}
