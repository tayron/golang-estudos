package mockfuncoes

import "testing"

func TestCalcularValorComMock(t *testing.T) {
	mockCalculadora := MockCalculadora{}
	result := CalcularValor(mockCalculadora)

	if result != 10 {
		t.Errorf("Resultado inesperado. Esperava-se 10, mas obteve %d", result)
	}
}

func TestCalcularValor(t *testing.T) {
	calculadora := Calculadora{}
	result := CalcularValor(calculadora)

	if result != 25 {
		t.Errorf("Resultado inesperado. Esperava-se 25, mas obteve %d", result)
	}
}
