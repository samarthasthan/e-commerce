package main

import (
	"fmt"

	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/internal/authentication/delivery/grpc"
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/kafka"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	tracer "github.com/samarthasthan/e-commerce/pkg/zipkin"
)

var (
	AUTHENTICATION_GRPC_PORT           string
	AUTHENTICATION_MYSQL_PORT          string
	AUTHENTICATION_MYSQL_ROOT_PASSWORD string
	AUTHENTICATION_MYSQL_HOST          string
	AUTHENTICATION_REDIS_PORT          string
	AUTHENTICATION_REDIS_HOST          string
	MAIL_GRPC_PORT                     string
	KAFKA_PORT                         string
	KAFKA_HOST                         string
)

func init() {
	AUTHENTICATION_GRPC_PORT = env.GetEnv("AUTHENTICATION_GRPC_PORT", "8000")
	AUTHENTICATION_MYSQL_PORT = env.GetEnv("AUTHENTICATION_MYSQL_PORT", "3306")
	AUTHENTICATION_MYSQL_ROOT_PASSWORD = env.GetEnv("AUTHENTICATION_MYSQL_ROOT_PASSWORD", "password")
	AUTHENTICATION_MYSQL_HOST = env.GetEnv("AUTHENTICATION_MYSQL_HOST", "localhost")
	AUTHENTICATION_REDIS_PORT = env.GetEnv("AUTHENTICATION_REDIS_PORT", "6379")
	AUTHENTICATION_REDIS_HOST = env.GetEnv("AUTHENTICATION_REDIS_HOST", "localhost")
	KAFKA_PORT = env.GetEnv("KAFKA_PORT", "9092")
	KAFKA_HOST = env.GetEnv("KAFKA_HOST", "localhost")
}

func main() {
	// Initialising Custom Logger
	log := logger.NewLogger("Authentication")

	// create a new Zipkin tracer for grpc server
	rt, rr, err := tracer.NewTracer("authentication", 8000)
	if err != nil {
		log.Fatalf("failed to create tracer: %v", err)
	}
	defer rr.Close()

	// create a new Zipkin tracer for mysql
	sqlt, sqlr, err := tracer.NewTracer("authentication-mysql", 8001)
	if err != nil {
		log.Fatalf("failed to create tracer: %v", err)
	}
	defer sqlr.Close()

	// Initialising Kafka Producer
	p := kafka.NewKafkaProducer(KAFKA_HOST, KAFKA_PORT)

	// Initialising Databases
	mysql := database.NewMySQL()
	redis := database.NewRedis()
	err = mysql.Connect(mysql.RegisterZipkin(sqlt), fmt.Sprintf("root:%s@tcp(%s:%s)/authentication?parseTime=true", AUTHENTICATION_MYSQL_ROOT_PASSWORD, AUTHENTICATION_MYSQL_HOST, AUTHENTICATION_MYSQL_PORT))
	if err != nil {
		log.Println(err.Error())
	}
	err = redis.Connect("redis", fmt.Sprintf("%s:%s", AUTHENTICATION_REDIS_HOST, AUTHENTICATION_REDIS_PORT))
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

	// Initialising gRPC Server
	server := grpc.NewAuthenticationGrpcServer(log, mysql, redis, p, rt)
	server.Run(AUTHENTICATION_GRPC_PORT)
}
