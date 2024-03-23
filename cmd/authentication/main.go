package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samarthasthan/e-commerce/internal/authentication"
	"github.com/samarthasthan/e-commerce/utils"
)

var (
	PORT            = utils.GetEnv("AUTH_PORT", "8001")
	ADDRESS         = fmt.Sprintf(":%s", PORT)
	MY_SQL_PASSWORD = utils.GetEnv("AUTH_MYSQL_ROOT_PASSWORD", "password")
	MY_SQL_PORT     = utils.GetEnv("AUTH_MYSQL_PORT", "1248")
	MY_SQL_HOST     = utils.GetEnv("AUTH_MYSQL_HOST", "localhost")
	MY_SQL_ADDRESS  = fmt.Sprintf("root:%s@(%s:%s)/", MY_SQL_PASSWORD, MY_SQL_HOST, MY_SQL_PORT)
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	s := authentication.MySql{}
	go func() {
		s.Connect(MY_SQL_ADDRESS)
		defer wg.Done()
	}()
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Fatalln(http.ListenAndServe(ADDRESS, mux))
	wg.Wait()
}
