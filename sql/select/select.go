package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:yakTLS&70c52@tcp(servidor_mysql_local:3306)/cursogo")

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 3)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
