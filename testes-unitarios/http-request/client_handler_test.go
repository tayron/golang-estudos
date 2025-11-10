package httprequest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/department?id=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	HelloWorldHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code, fmt.Sprintf("Esperado status %d mas foi retornado %d", http.StatusOK, w.Code))

	valorEsperado := `{"message":"Hello, World!"}`

	assert.JSONEq(t, valorEsperado, w.Body.String(),
		fmt.Sprintf("Esperado conteúdo: '%s' mas foi retornado '%s'", valorEsperado, w.Body.String()))
}
