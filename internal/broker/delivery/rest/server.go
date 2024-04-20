package rest

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samarthasthan/e-commerce/internal/broker/delivery/grpc"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

var (
	wg             sync.WaitGroup
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests by method and path.",
		},
		[]string{"method", "path"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time in seconds by method and path.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}

// MetricsMiddleware is a middleware that tracks request count and duration.
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime).Seconds()
		requestCounter.WithLabelValues(r.Method, r.URL.Path).Inc()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}

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

	wrappedServeMux := MetricsMiddleware(r.ServeMux)
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), wrappedServeMux)
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

	r.ServeMux.Handle("/metrics", promhttp.Handler())
}
