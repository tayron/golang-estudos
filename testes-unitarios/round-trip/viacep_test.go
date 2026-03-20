package roundtrip_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	roundtrip "github.com/tayron/golang-estudos/testes-unitarios/round-trip"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestSearchCEP(t *testing.T) {
	cep := "01001000"

	t.Run("Teste consulta com sucesso", func(t *testing.T) {
		client := &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {

				// 🔍 você pode validar a request
				url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
				if req.URL.String() != url {
					t.Fatalf("url incorreta: %s", req.URL.String())
				}

				// 📦 resposta fake
				body := `{
					"cep": "01001-000",
					"logradouro": "Praça da Sé",
					"complemento": "lado ímpar",
					"unidade": "",
					"bairro": "Sé",
					"localidade": "São Paulo",
					"uf": "SP",
					"estado": "São Paulo",
					"regiao": "Sudeste",
					"ibge": "3550308",
					"gia": "1004",
					"ddd": "11",
					"siafi": "7107"
				}`

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(body)),
				}, nil
			}),
		}

		address, err := roundtrip.SearchCEP(client, cep)
		if err != nil {
			t.Fatalf("expected error nil but got: '%s'", err.Error())
		}

		if address.Logradouro != "Praça da Sé" {
			t.Errorf("Expected logradouro 'Praça da Sé' but got: '%s'", address.Logradouro)
		}
	})

}
