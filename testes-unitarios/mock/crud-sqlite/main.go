package main

import (
	"github.com/tayron/golang-estudos/testes-unitarios/mock/crud-sqlite/repository"

	"github.com/chromedp/cdproto/database"
)

func main() {
	repository.CreateUser(database.Database, "Pedro", 37)
}
