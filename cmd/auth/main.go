package main

import (
	"log"

	"github.com/samarthasthan/e-commerce-backend/internal/auth"
)

func main() {
	s := auth.NewAuthSqlDB()
	err := s.Migrate()
	if err != nil {
		log.Fatalln(err)
	}
	defer s.DB.Close()
}
