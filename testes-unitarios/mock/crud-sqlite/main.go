package main

import (
	"example.com/mock/crud-sqlite/repository"

	"github.com/chromedp/cdproto/database"
)

func main() {
	repository.CreateUser(database.Database, "Pedro", 37)
}
