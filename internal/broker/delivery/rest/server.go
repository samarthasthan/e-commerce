package rest

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/samarthasthan/e-commerce/internal/broker/delivery/grpc"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

var (
	wg sync.WaitGroup
)

type RestServer struct {
	ServeMux *http.ServeMux
	log      *logger.Logger
	handler  *Handler
}

func NewRestServer(log *logger.Logger) *RestServer {
	g := grpc.NewGRPCClients(log)
	g.ConnectToAuthenticationService()
	h := NewHandler(log, g)
	m := http.NewServeMux()
	return &RestServer{
		ServeMux: m,
		log:      log,
		handler:  h,
	}
}

func (r *RestServer) RunServer(PORT string) {
	r.HandleRoutes()
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), r.ServeMux)
		if err != nil {
			r.log.Errorln(err)
		}
		defer wg.Done()
	}()
	r.log.Infof("Broker listing on: %v", PORT)
	wg.Wait()
}

func (r *RestServer) HandleRoutes() {
	r.ServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello welcome to broker"))
	})
	r.ServeMux.Handle("POST /signup", LoggingMiddleware(r.log, http.HandlerFunc(r.handler.SignUp)))
}
