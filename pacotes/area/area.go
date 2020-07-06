package area

import "math"

// Pi é uma proporção numérica definida pela relção entre
// o perímetro de uma circuferência e seu diâmetro
const Pi = 3.1416

// Circ é repsonsável por calcular a área da circuferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsaǘel por clacular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TringuloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
