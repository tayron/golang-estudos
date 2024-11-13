package main

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

// Modelo de dados do usuário com validações
type User struct {
	FullName        string `json:"full_name" validate:"required"`                         // Nome completo (obrigatório)
	Email           string `json:"email" validate:"required,email"`                       // Email (obrigatório e formato válido)
	Password        string `json:"password" validate:"required,min=6"`                    // Senha (obrigatória, mínimo de 6 caracteres)
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"` // Confirmar senha (obrigatório, igual à senha)
	CreatedAt       int64  `json:"created_at"`                                            // Data de criação do usuário
	UpdatedAt       int64  `json:"updated_at"`                                            // Data de atualização do usuário
}

// Função para criar mensagens de erro personalizadas
func getErrorMessages(err validator.ValidationErrors) map[string]string {
	// Mapeamento para mensagens de erro personalizadas
	messages := map[string]string{}

	for _, fieldErr := range err {
		var message string
		switch fieldErr.Field() {
		case "FullName":
			message = "Nome completo do usuário deve ser informado."
		case "Email":
			message = "Email deve ser informado e estar em um formato válido."
		case "Password":
			if fieldErr.Tag() == "min" {
				message = "A senha deve ter no mínimo 6 caracteres."
			} else {
				message = "A senha é obrigatória."
			}
		case "ConfirmPassword":
			message = "A confirmação de senha deve ser igual à senha."
		default:
			message = "Erro no campo " + fieldErr.Field()
		}
		messages[fieldErr.Field()] = message
	}
	return messages
}

// Função para criar um novo usuário e validar os dados
func createUser() {
	// Exemplo de dados do usuário
	user := User{
		FullName:        "", // Nome completo vazio para gerar erro de validação
		Email:           "john.doe@example.com",
		Password:        "pass",
		ConfirmPassword: "pass",
		CreatedAt:       time.Now().Unix(),
		UpdatedAt:       time.Now().Unix(),
	}

	// Instância do validador
	validate := validator.New()

	// Validação dos dados do usuário
	err := validate.Struct(user)
	if err != nil {
		// Obter mensagens de erro personalizadas
		validationErrors := getErrorMessages(err.(validator.ValidationErrors))
		for field, message := range validationErrors {
			fmt.Printf("Erro no campo '%s': %s\n", field, message)
		}
	} else {
		fmt.Println("Dados do usuário são válidos:", user)
	}
}

func main() {
	createUser()
}
