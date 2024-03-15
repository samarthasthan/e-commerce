package broker

import (
	"net/http"
)

type BrokerServer struct {
	mux *http.ServeMux
}

func NewBrokerServer() *BrokerServer {
	mux := http.NewServeMux()
	return &BrokerServer{mux: mux}
}

func (s *BrokerServer) Run(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func (s *BrokerServer) HandleRoutes() {
	s.mux.HandleFunc("/", welcomeHandler)
	s.HandleAuthRoutes() // Handle authentication or authorization related routes
}

func (s *BrokerServer) HandleAuthRoutes() {
	s.mux.HandleFunc("POST /signup", signUp)
	// s.mux.HandleFunc("POST /checkphone",)
}
