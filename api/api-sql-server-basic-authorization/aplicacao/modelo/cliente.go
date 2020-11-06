package modelo

import (
	"fmt"

	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/biblioteca/banco"
	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/biblioteca/util"
	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/entidade"
)

// ObterClientePorCPF -
func ObterClientePorCPF(CpfCnpj string) entidade.Cliente {
	db := banco.ObterConexaoDW()
	defer db.Close()

	cliente := entidade.Cliente{}

	var sql string = `SELECT ID, NOME, CPF FROM CLIENTES WHERE CPF = '%s'`

	queryString := fmt.Sprintf(sql, CpfCnpj)
	rows, err := db.Query(queryString)

	if err != nil {
		util.GravarLog("/aplicacao/modelo/cliente.go::ObterClientePorCPF", err.Error())
		return cliente
	}

	defer rows.Close()

	for rows.Next() {
		errScan := rows.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.CPF,
		)

		if errScan != nil {
			util.GravarLog("/aplicacao/modelo/cliente.go::ObterClientePorCPF", errScan.Error())
		}

		return cliente
	}

	return cliente
}

// ObterTodosClientes -
func ObterTodosClientes(pagina string, numeroRegistros string) (listaClientes []entidade.Cliente) {
	db := banco.ObterConexaoDW()
	defer db.Close()

	var sql string = `SELECT ID, NOME, CPF FROM DIM_CLIENTES
	ORDER BY ID
	OFFSET ((%s - 1) * %s) ROWS
	FETCH NEXT %s ROWS ONLY`

	queryString := fmt.Sprintf(sql, pagina, numeroRegistros, numeroRegistros)
	rows, err := db.Query(queryString)

	if err != nil {
		util.GravarLog("/aplicacao/modelo/cliente.go::ObterTodosClientes", err.Error())
		return listaClientes
	}

	defer rows.Close()

	for rows.Next() {

		var cliente entidade.Cliente

		errScan := rows.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.CPF,
		)

		if errScan != nil {
			util.GravarLog("/aplicacao/modelo/cliente.go::ObterTodosClientes", errScan.Error())
		}

		listaClientes = append(listaClientes, cliente)
	}

	return listaClientes
}
