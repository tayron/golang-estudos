package mockfuncoes

func CalcularValor(calculadora CalculadoraInterface) int {
	return calculadora.Somar(10, 15)
}
