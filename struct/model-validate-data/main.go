package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptTranslations "github.com/go-playground/validator/v10/translations/pt"
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
	Vencimento         string  `json:"vencimento" validate:"required,datetime=02/01/2006"`
}

type RequestValidator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func NewRequestValidator() (*RequestValidator, error) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		jsonName := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonName == "" || jsonName == "-" {
			return field.Name
		}
		return jsonName
	})

	ptLocale := pt.New()
	uni := ut.New(ptLocale, ptLocale)
	translator, found := uni.GetTranslator("pt")
	if !found {
		return nil, fmt.Errorf("tradutor pt nao encontrado")
	}

	if err := ptTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
		return nil, err
	}

	return &RequestValidator{
		validate:   validate,
		translator: translator,
	}, nil
}

func (rv *RequestValidator) Validate(data RequestData) []string {
	err := rv.validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}

	errors := make([]string, 0, len(validationErrors))
	for _, validationErr := range validationErrors {
		errors = append(errors, validationErr.Translate(rv.translator))
	}

	return errors
}

func main() {
	validatorInstance, err := NewRequestValidator()
	if err != nil {
		panic(err)
	}

	requests := []RequestData{
		{
			CnpjAdministradora: "12345678000195",
			CnpjRemetente:      "12345678000195",
			CnpjFornecedor:     "12345678000195",
			CnpjCondominio:     "12345678000195",
			TipoInscrCondo:     "J",
			TipoInscrFornec:    "F",
			Ano:                2024,
			Mes:                11,
			RazaoSocial:        "ZINCA",
			Produto:            "Nome do produto",
			Valor:              3699.55,
			Vencimento:         "30/12/2024",
		},
		{
			CnpjAdministradora: "123",
			CnpjRemetente:      "abcdefghijklmn",
			CnpjFornecedor:     "",
			CnpjCondominio:     "12345678000195",
			TipoInscrCondo:     "X",
			TipoInscrFornec:    "J",
			Ano:                1800,
			Mes:                13,
			RazaoSocial:        "",
			Produto:            "Nome do produto",
			Valor:              0,
			Vencimento:         "12/30/2024",
		},
	}

	for index, data := range requests {
		fmt.Printf("Validando payload %d\n", index+1)

		errors := validatorInstance.Validate(data)
		if len(errors) == 0 {
			fmt.Println("dados validados com sucesso")
			fmt.Println()
			continue
		}

		fmt.Println("erros encontrados:")
		for _, validationError := range errors {
			fmt.Printf("- %s\n", validationError)
		}
		fmt.Println()
	}
}
