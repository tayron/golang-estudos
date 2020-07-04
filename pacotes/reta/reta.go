package main

import "math"

// Variavel e métodos
// Iniciando com letra maiúscula é PÚBLICO (visível dentro e fora do pacote)
// Iniciando com letra minúscula é PRIVADO (visível apenas dentro do pacote)

// Por exemplo...
// Ponto - gerará algo público
// Ponto ou _ponto - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsavel por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
