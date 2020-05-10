package database

import (
	"database/sql"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// CreateDatabaseConnection makes a connection into database
func CreateDatabaseConnection() *sql.DB {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/apiGo")
	if err != nil {
		panic(err.Error())
	}

	return db
}
