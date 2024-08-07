package rest

import (
	"net/http"
	"time"

	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samarthasthan/e-commerce/internal/broker/validation"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/pkg/proto_go"
)

var (
	// Request counters
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Number of requests received",
		},
		[]string{"path"},
	)

	// Request duration histograms
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

type RestHandler struct {
	mux                  *http.ServeMux
	authenticationClient proto_go.AuthenticationServiceClient
	validator            *validation.Validator
	log                  *logger.Logger
	tracer               *zipkin.Tracer
}

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}

func NewRestHandler(ac proto_go.AuthenticationServiceClient, v *validation.Validator, l *logger.Logger, m *http.ServeMux, t *zipkin.Tracer) *RestHandler {
	return &RestHandler{
		mux:                  m,
		authenticationClient: ac,
		validator:            v,
		log:                  l,
		tracer:               t,
	}
}

func (s *RestHandler) Handle() {
	// create a middleware that traces incoming requests
	zipkinMiddleWare := zipkinhttp.NewServerMiddleware(
		s.tracer,
		zipkinhttp.SpanName("broker"),
	)

	s.mux.Handle("/metrics", promhttp.Handler())
	s.mux.Handle("POST /auth/register", MetricsMiddleware(zipkinMiddleWare(http.HandlerFunc(s.CreateUser))))
	s.mux.Handle("POST /auth/otp-verify", MetricsMiddleware(zipkinMiddleWare(http.HandlerFunc(s.OTPVerify))))
	s.mux.Handle("POST /auth/login", MetricsMiddleware(zipkinMiddleWare(http.HandlerFunc(s.SignInUser))))
}

func (s *RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path

		// Increment the request counter
		requestCounter.WithLabelValues(path).Inc()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Observe the request duration
		duration := float64(time.Since(start).Milliseconds())
		requestDuration.WithLabelValues(path).Observe(duration)
	})
}
