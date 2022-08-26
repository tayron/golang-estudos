package models

import (
	"os"

	"github.com/tayron/golang-estudos/paginacao/application/database"
)

type Produto struct {
	ID         int
	Referencia string
	Cor        string
	Tamanho    string
}

// BuscarTodosProdutos - Retorna todos os produtos
func BuscarTodosProdutos(offset int) []Produto {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT id, referencia, cor, tamanho FROM produtos ORDER BY id DESC LIMIT ? OFFSET ?`

	numeroRegistro := os.Getenv("NUMERO_REGISTRO_POR_PAGINA")
	rows, err := db.Query(sql, numeroRegistro, offset)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var listaProdutos []Produto
	for rows.Next() {

		var produtoStruct Produto

		rows.Scan(&produtoStruct.ID,
			&produtoStruct.Referencia,
			&produtoStruct.Cor,
			&produtoStruct.Tamanho)

		listaProdutos = append(listaProdutos, produtoStruct)
	}

	return listaProdutos
}

// ObterNumeroProdutos - Retorna número de produtos
func ObterNumeroProdutos() int {

	db := database.ObterConexao()
	defer db.Close()

	var sql string = `SELECT count(0) FROM produtos`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var numeroProdutos int = 0
	for rows.Next() {
		rows.Scan(&numeroProdutos)
	}

	return numeroProdutos
}
