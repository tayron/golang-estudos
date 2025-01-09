package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt"
)

type RequestData struct {
	CnpjAdministradora string  `json:"cnpj_administradora" validate:"required,len=14,numeric"`
	CnpjRemetente      string  `json:"cnpj_remetente" validate:"required,len=14,numeric"`
	CnpjFornecedor     string  `json:"cnpj_fornecedor" validate:"required,len=14,numeric"`
	CnpjCondominio     string  `json:"cnpj_condominio" validate:"required,len=14,numeric"`
	TipoInscrCondo     string  `json:"tipo_inscr_condo" validate:"required,oneof=J F"`
	TipoInscrFornec    string  `json:"tipo_inscr_fornec" validate:"required,oneof=J F"`
	Ano                int     `json:"ano" validate:"required,min=1900,max=2100"`
	Mes                int     `json:"mes" validate:"required,min=1,max=12"`
	RazaoSocial        string  `json:"razao_social" validate:"required,max=100"`
	Produto            string  `json:"produto" validate:"required,max=100"`
	Valor              float64 `json:"valor" validate:"required,gt=0"`
	Vencimento         string  `json:"vencimento" validate:"required,datetime=01/02/2006"`
}

func (r RequestData) validateData() error {
	validate := validator.New()

	ptLocale := pt.New()
	uni := ut.New(ptLocale, ptLocale)
	translator, _ := uni.GetTranslator("pt")
	_ = pt_translations.RegisterDefaultTranslations(validate, translator)

	var validationErrors []string
	err := validate.Struct(r)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Translate(translator))
		}
	}

	if len(validationErrors) > 0 {
		return errors.New("Erros de validação: " + strings.Join(validationErrors, ", "))
	}
	return nil
}

func main() {
	requestData := []RequestData{
		{
			CnpjAdministradora: "12345678000195",
			CnpjRemetente:      "12345678000195",
			CnpjFornecedor:     "12345678000195",
			CnpjCondominio:     "12345678000195",
			TipoInscrCondo:     "J",
			TipoInscrFornec:    "J",
			Ano:                2019,
			Mes:                11,
			RazaoSocial:        "ZINCA",
			Produto:            "Nome do produto",
			Valor:              3699.55,
			Vencimento:         "12/30/2018",
		},
	}

	for _, data := range requestData {
		if err := data.validateData(); err != nil {
			fmt.Printf("Erro de validação: %v\n", err)
			return
		}
	}

	fmt.Println("Dados validados com sucesso!")
}
