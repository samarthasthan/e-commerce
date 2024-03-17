package authentication

import (
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

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
			log.Info("failed to fetch URL",
				// Structured context as strongly typed Field values.
				zap.String("url", "gogoel.com"),
				zap.Int("attempt", 3),
				zap.Duration("backoff", time.Second),
			)
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
