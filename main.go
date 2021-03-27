package main

import (
	"gin_test/db"
	"gin_test/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.Open()
	db.Init()
	defer db.Close()
	router.Router()
}
