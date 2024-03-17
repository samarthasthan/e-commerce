package broker

import (
	"log"
	"net/http"
)

type BrokerServer struct {
	S *http.ServeMux
}

func NewBrokerServer(addr string) {
	mux := http.NewServeMux()

	log.Printf("Broker listening on PORT %s", addr)
	log.Fatalln(http.ListenAndServe(addr, mux))
}
