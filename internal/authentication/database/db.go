package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	zipkinsql "github.com/openzipkin-contrib/zipkin-go-sql"
	"github.com/openzipkin/zipkin-go"
	"github.com/redis/go-redis/v9"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
)

type Database interface {
	Connect(string,string) error
	Close() error
	RegisterZipkin(*zipkin.Tracer) string
}

type MySQL struct {
	Queries *sqlc.Queries
	DB      *sql.DB
}

type Redis struct {
	RDB *redis.Client
}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func NewRedis() *Redis {
	return &Redis{}
}

func (s *MySQL) Connect(driverName string, addr string) error {
	db, err := sql.Open(driverName, addr)
	if err != nil {
		return err
	}
	s.DB = db
	s.Queries = sqlc.New(db)
	return nil
}

func (s *MySQL) RegisterZipkin(tracer *zipkin.Tracer) string {
	// Register our zipkinsql wrapper for the provided MySQL driver.
	driverName, err := zipkinsql.Register("mysql", tracer, zipkinsql.WithAllTraceOptions())
	if err != nil {
		log.Fatalf("unable to register zipkin driver: %v\n", err)
	}
	return driverName
}

func (s *MySQL) Close() error {
	err := s.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Connect(string,addr string) error {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	r.RDB = rdb
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}

func (r *Redis) RegisterZipkin(tracer *zipkin.Tracer) string {
	return "redis"
}

func (r *Redis) Close() error {
	err := r.RDB.Close()
	if err != nil {
		return err
	}
	return nil
}
