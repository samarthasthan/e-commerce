package main

import (
	"fmt"
	"log"

	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/pkg/env"
)

var (
	AUTHENTICATION_GRPC_PORT           string
	AUTHENTICATION_MYSQL_PORT          string
	AUTHENTICATION_MYSQL_ROOT_PASSWORD string
	AUTHENTICATION_MYSQL_HOST          string
	AUTHENTICATION_REDIS_PORT          string
	AUTHENTICATION_REDIS_HOST          string
)

func init() {
	AUTHENTICATION_GRPC_PORT = env.GetEnv("AUTHENTICATION_GRPC_PORT", "8000")
	AUTHENTICATION_MYSQL_PORT = env.GetEnv("AUTHENTICATION_MYSQL_PORT", "8001")
	AUTHENTICATION_MYSQL_ROOT_PASSWORD = env.GetEnv("AUTHENTICATION_MYSQL_ROOT_PASSWORD", "password")
	AUTHENTICATION_MYSQL_HOST = env.GetEnv("AUTHENTICATION_MYSQL_HOST", "localhost")
	AUTHENTICATION_REDIS_PORT = env.GetEnv("AUTHENTICATION_REDIS_PORT", "8002")
	AUTHENTICATION_REDIS_HOST = env.GetEnv("AUTHENTICATION_REDIS_HOST", "localhost")
}

func main() {
	// Initialising Databases
	mysql := database.NewMySQL()
	redis := database.NewRedis()
	err := mysql.Connect(fmt.Sprintf("root:%s@tcp(%s:%s)/authentication", AUTHENTICATION_MYSQL_ROOT_PASSWORD, AUTHENTICATION_MYSQL_HOST, AUTHENTICATION_MYSQL_PORT))
	if err != nil {
		log.Println(err.Error())
	}
	err = redis.Connect(fmt.Sprintf("%s:%s", AUTHENTICATION_REDIS_HOST, AUTHENTICATION_REDIS_PORT))
	if err != nil {
		log.Println(err.Error())
	}
	defer func() {
		err = mysql.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	defer func() {
		err = redis.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()
}
