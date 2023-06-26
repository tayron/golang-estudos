package method

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloYou(t *testing.T) {
	p := Person{
		Name: "Tayron",
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(p)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello-you", &b)
	HelloYou(rr, req)

	if http.StatusOK != rr.Code {
		t.Errorf("Status code esperado: %d, status recuperado: %d", http.StatusOK, rr.Code)
	}

	response, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Error(err)
	}

	expect := fmt.Sprintf("Hello %s\n", p.Name)

	if string(response) != expect {
		t.Errorf("Mensagem esperada: %s, status recuperado: %s", expect, string(response))
	}
}
