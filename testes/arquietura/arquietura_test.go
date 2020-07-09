package arquietura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // Teste em paralelo
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquietura amd64")
	}

	t.Fail()
}
