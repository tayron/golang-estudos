package entity

type CalculadoraMock struct {
	Resultado int
}

func (c *CalculadoraMock) Somar(a int, b int) {
	c.Resultado = 5
}
