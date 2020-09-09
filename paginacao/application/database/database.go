package database

import (
	"database/sql"
	"fmt"
	"os"

	// Uso indireto do drive mysql
	_ "github.com/go-sql-driver/mysql"
)

// ExecutarQuery - Executa query no banco de dados
func ExecutarQuery(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

// ObterConexao - retorna conexão com banco de dados
func ObterConexao() *sql.DB {

	stringConexao := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_LOCALHOST"),
		os.Getenv("DB_PORTA"),
	)

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		panic(err)
	}

	ExecutarQuery(db, "create database if not exists "+os.Getenv("DB_BANCO"))
	ExecutarQuery(db, "use "+os.Getenv("DB_BANCO"))

	return db
}
