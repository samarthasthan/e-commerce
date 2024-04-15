package broker

import (
	"fmt"
	"net/http"
	"sync"

	constants "github.com/samarthasthan/e-commerce"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

type RestServer struct {
	ServeMux *http.ServeMux
	Logger   *logrus.Logger
}

func NewRestServer() *RestServer {
	l := logger.NewLogger(constants.BROKER_LOGGER_NAME)
	m := http.NewServeMux()
	return &RestServer{
		Logger:   l.Logger,
		ServeMux: m,
	}
}

func (r *RestServer) RunServer(PORT string) {
	r.HandleRoutes()
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), r.ServeMux)
		if err != nil {
			r.Logger.Errorln(err)
		}
		wg.Done()
	}()
	r.Logger.Infof("Broker listing on: %v", PORT)
	wg.Wait()
}

func (r *RestServer) HandleRoutes() {
	r.ServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello welcome to broker"))
	})
}
