package broker

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type BrokerServer struct {
	S *http.ServeMux
}

func NewBrokerServer(port string) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Printf("Broker listening on PORT %s", port)
	log.Fatalln(http.ListenAndServe(port, mux))
}
