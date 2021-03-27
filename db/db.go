package db

import (
	"database/sql"
	"log"
)

var db *sql.DB

const DB_SOURCE_NAME = "./db/sample.db"
const DB_DRIVER_NAME = "sqlite3"

func createTable() {
	q := `
		CREATE TABLE todo (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(225) NOT NULL,
			done INTEGER NOT NULL DEFAULT 0,
			created TIMESTAMP DEFAULT (DATETIME('now', 'localtime')),
			updated TIMESTAMP DEFAULT (DATETIME('now', 'localtime'))
		)
	`

	statement, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("todo table created")
}

func deleteTable() {
	q := `
		DROP TABLE IF EXISTS todo
	`

	statement, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("todo table created")
}

func addInitialTableData() {
	q := `
		INSERT INTO todo(name, done)
		VALUES ('sample todo 1', false), ('sample todo 2', true)
	`
	statement, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("add initial data")
}

func Get() *sql.DB {
	return db
}

func Open() {
	_db, err := sql.Open(DB_DRIVER_NAME, DB_SOURCE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	db = _db
}

func Init() {
	deleteTable()
	createTable()
	addInitialTableData()
}

func Close() {
	db.Close()
}
