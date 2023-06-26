package method

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello-world", nil)
	HelloWorld(rr, req)

	if http.StatusOK != rr.Code {
		t.Errorf("Status code esperado: %d, status recuperado: %d", http.StatusOK, rr.Code)
	}
}
