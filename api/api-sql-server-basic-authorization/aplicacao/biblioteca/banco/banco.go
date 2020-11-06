package banco

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	// Uso indireto do drive mysql
	"github.com/tayron/go-lang-estudos/api/api-sql-server-basic-authorization/aplicacao/biblioteca/util"
	_ "github.com/denisenkom/go-mssqldb"
)

// ExecutarQuery - Executa query no banco de dados
func ExecutarQuery(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

// ObterConexaoDW - retorna conexão com banco de dados
func ObterConexaoDW() *sql.DB {
	query := url.Values{}
	query.Add("database", os.Getenv("DB_DW_BANCO"))

	u := url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(os.Getenv("DB_DW_USUARIO"), os.Getenv("DB_DW_SENHA")),
		Host:   fmt.Sprintf("%s:%s", os.Getenv("DB_DW_HOST"), os.Getenv("DB_DW_PORTA")),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())

	if err != nil {
		util.GravarLog("/aplicacao/biblioteca/banco/banco.go::ObterConexaoDW", err.Error())
		panic(err)
	}

	return db
}
