package auth

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type AuthDB interface{}

type SqlDB struct {
	DB *sqlx.DB
}

func NewAuthSqlDB() *SqlDB {
	db, err := sqlx.Connect("mysql", "root:password@tcp(localhost:3306)/")
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &SqlDB{DB: db}
}
