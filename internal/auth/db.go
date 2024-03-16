package auth

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jmoiron/sqlx"
)

type AuthDB interface{}

type SqlDB struct {
	DB *sqlx.DB
}

type MongoDB struct {
	DB *mongo.Client
}

type Redis struct {
	DB *redis.Client
}

func NewAuthSqlDB() *SqlDB {
	db, err := sqlx.Connect("mysql", "root:password@tcp(localhost:1248)/")
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &SqlDB{DB: db}
}

// func NewAuthMongoDB() *MongoDB {

// }
