package database

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
)

type Database interface {
	Connect(string) error
	Close() error
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

func (s *MySQL) Connect(addr string) error {
	db, err := sql.Open("mysql", addr)
	if err != nil {
		return err
	}
	s.DB = db
	s.Queries = sqlc.New(db)
	return nil
}

func (s *MySQL) Close() error {
	err := s.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Connect(addr string) error {
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

func (r *Redis) Close() error {
	err := r.RDB.Close()
	if err != nil {
		return err
	}
	return nil
}
