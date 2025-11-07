package mockfuncoes

type Calculadora struct{}

func (m Calculadora) Somar(a, b int) int {
	return a + b
}
