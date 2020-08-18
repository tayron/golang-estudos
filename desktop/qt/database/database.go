package database

import (
	"database/sql"
	"fmt"
	"os"

	// Uso indireto do drive mysql
	_ "github.com/go-sql-driver/mysql"
)

// ObterConexao - retorna conexão com banco de dados
func ObterConexao() *sql.DB {

	stringConexao := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_LOCALHOST"),
		os.Getenv("DB_PORTA"),
		os.Getenv("DB_BANCO"),
	)

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		panic(err)
	}

	return db
}
