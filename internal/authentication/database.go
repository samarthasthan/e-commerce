package authentication

import (
	"github.com/jmoiron/sqlx"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type MySql struct {
	DB *sqlx.DB
}

type Mongo struct {
	DB *mongo.Client
}

type Redis struct {
	DB *redis.Client
}

func (s *MySql) Connect(addr string) *MySql {
	d, err := sqlx.Connect("mysql", addr)
	if err != nil {
		for {
			log.Error(err)
		}
	}
	return &MySql{
		DB: d,
	}
}

func (m *Mongo) Connect(addr string) *Mongo {
	return &Mongo{}
}

func (r *Redis) Connect(addr string) *Redis {
	return &Redis{}
}
