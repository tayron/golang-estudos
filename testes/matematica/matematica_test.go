package matematica

import "testing"

const erroPadrao = "O valor esperado %v, mas o resultado encontrado foi %v."

// comando "go test" dentro do diretório onde o teste se encontra
// "go test ./..." na raiz para executar todos os tests
func TestMedia(t *testing.T) {
	t.Parallel() // Teste em paralelo
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
