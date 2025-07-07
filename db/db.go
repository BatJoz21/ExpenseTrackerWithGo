package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DataBase *sql.DB

func InitDB() {
	var err error

	DataBase, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("failed to connect to database")
	}

	DataBase.SetMaxOpenConns(10)
	DataBase.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	queryUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullname TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL
	)`

	_, err := DataBase.Exec(queryUserTable)
	if err != nil {
		panic("failed to create users table")
	}

	queryExpenseTable := `CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		account TEXT NOT NULL,
		amount INTEGER NOT NULL,
		category TEXT NOT NULL,
		date DATETIME NOT NULL,
		expense_type TEXT NOT NULL,
		note TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DataBase.Exec(queryExpenseTable)
	if err != nil {
		panic("failed to create expense table")
	}
}
