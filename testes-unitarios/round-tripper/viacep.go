package roundtripper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func sanitizeCEP(cep string) string {
	return strings.ReplaceAll(cep, "-", "")
}

func SearchCEP(client *http.Client, cep string) (*ViaCEPResponse, error) {
	if client == nil {
		client = &http.Client{Timeout: 5 * time.Second}
	}

	cep = sanitizeCEP(cep)
	if cep == "" {
		return nil, fmt.Errorf("invalid cep")
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var result ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// ViaCEP retorna campo "erro" quando não encontra
	if result.CEP == "" {
		return nil, fmt.Errorf("cep not found")
	}

	return &result, nil
}
