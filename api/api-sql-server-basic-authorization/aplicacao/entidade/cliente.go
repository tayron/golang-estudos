package entidade

import (
	"github.com/tayron/golang-estudos/api/api-sql-server-basic-authorization/aplicacao/biblioteca/util"
)

type Cliente struct {
	ID   int
	Nome string
	CPF  *string
}

// GetNome -
func (c Cliente) GetNome() string {
	return util.ConverteStringUperCase(c.Nome)
}

// GetCPF -
func (c Cliente) GetCPF() string {
	if c.CPF != nil {
		return util.RemoverCaracteresEspeciais(*c.CPF)
	}

	return "9"
}
