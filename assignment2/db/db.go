package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@/my_database")
	if err != nil {
		return nil, err
	}
	return db, nil
}
