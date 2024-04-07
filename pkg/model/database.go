package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

// InitDB initializes and returns a database connection
func InitDb(dsn string) (*DB, error) {
	// Open a db connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// verify the connection by pinging the database
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// return the database object
	return &DB{db}, nil
}
