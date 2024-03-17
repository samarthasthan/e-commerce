package main

import "github.com/samarthasthan/e-commerce/internal/authentication"

func main() {
	s := authentication.MySql{}
	s.Connect("root:password@tcp(localhost:1248)/")
}
