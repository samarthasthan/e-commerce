package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var port = ":8004"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Printf("Broker listening on PORT %s", port)
	log.Fatalln(http.ListenAndServe(port, mux))
}
