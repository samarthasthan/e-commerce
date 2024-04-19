package broker

import (
	"fmt"
	"net/http"
	"sync"

	constants "github.com/samarthasthan/e-commerce"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

var (
	logrs = logger.NewLogger(constants.BROKER_LOGGER_NAME)
	wg    sync.WaitGroup
)

type RestServer struct {
	ServeMux *http.ServeMux
}

func NewRestServer() *RestServer {
	m := http.NewServeMux()
	return &RestServer{
		ServeMux: m,
	}
}

func (r *RestServer) RunServer(PORT string) {
	r.HandleRoutes()
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), r.ServeMux)
		if err != nil {
			logrs.Errorln(err)
		}
		wg.Done()
	}()
	logrs.Infof("Broker listing on: %v", PORT)
	wg.Wait()
}

func (r *RestServer) HandleRoutes() {
	r.ServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello welcome to broker"))
	})
	r.ServeMux.HandleFunc("POST /signup", SignUp)
}
