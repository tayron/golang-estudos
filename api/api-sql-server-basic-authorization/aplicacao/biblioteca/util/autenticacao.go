package util

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"strings"
)

// ValidarPermissaoAcesso -
func ValidarPermissaoAcesso(r *http.Request) error {
	autorizacao := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(autorizacao) != 2 || autorizacao[0] != "Basic" {
		return errors.New("Acesso negado")
	}

	payload, _ := base64.StdEncoding.DecodeString(autorizacao[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !validarUsuarioSenha(pair[0], pair[1]) {
		return errors.New("Acesso negado")
	}

	return nil
}

func validarUsuarioSenha(username, password string) bool {
	if username == os.Getenv("USUARIO") && password == os.Getenv("SENHA") {
		return true
	}
	return false
}
