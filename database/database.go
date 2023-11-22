package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	dsn := "root@tcp(localhost:3306)/golang-todos"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}