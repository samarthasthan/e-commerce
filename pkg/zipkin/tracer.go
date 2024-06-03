package tracer

import (
	"fmt"
	"log"
	"net"

	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/samarthasthan/e-commerce/pkg/env"
)

var (
	ZIPKIN_HOST string
	ZIPKIN_PORT string
)

func init() {
	ZIPKIN_HOST = env.GetEnv("ZIPKIN_HOST", "localhost")
	ZIPKIN_PORT = env.GetEnv("ZIPKIN_PORT", "9411")
}

func NewTracer(serviceName string, port uint16) (*zipkin.Tracer, error) {

	// Create a reporter to be used by the tracer
	reporter := httpreporter.NewReporter(fmt.Sprintf("http://%s:%s/api/v2/spans", ZIPKIN_HOST, ZIPKIN_PORT))

	// Local endpoint represent the local service information
	localEndpoint := &model.Endpoint{ServiceName: serviceName, IPv4: getOutboundIP(), Port: port}

	// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00)
	// of traces.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, err
	}

	tracer, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return nil, err
	}

	return tracer, err
}

// Get preferred outbound ip of this machine
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
