package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver mySql
)

// Connect open database connection and return
func Connect() (*sql.DB, error) {

	db, error := sql.Open("mysql", config.ConnectionString)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}
