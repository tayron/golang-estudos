package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:yakTLS&70c52@tcp(servidor_mysql_local:3306)/cursogo")

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Marcelo", 1)
	stmt.Exec("Patrícia", 2)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)

}
