package main

import (
	"fmt"
	"log"

	"github.com/samarthasthan/e-commerce-backend/internal/auth"
)

func init() {
	fmt.Println("Auth service started")
}

func main() {
	s := auth.NewAuthSqlDB()
	err := s.Migrate()
	if err != nil {
		log.Fatalln(err)
	}
	defer s.DB.Close()
}
