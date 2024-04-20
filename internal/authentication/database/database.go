package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
)

type Database struct {
	Queries *sqlc.Queries
	DB      *sql.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &Database{
		Queries: sqlc.New(db),
	}, nil
}

func (db *Database) Close() {
	err := db.DB.Close()
	if err != nil {
		log.Println("Error closing database:", err)
	}
}
