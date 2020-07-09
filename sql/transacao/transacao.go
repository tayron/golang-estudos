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

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios (id, nome) values (?, ?)")

	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Thiago")

	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	}

	tx.Commit()

}
