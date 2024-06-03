package main

import (
	"log"
	"time"

	"github.com/openzipkin/zipkin-go"
	zipkinc "github.com/samarthasthan/e-commerce/pkg/zipkin"
)

func main() {
	// create a new Zipkin tracer
	tracer, reporter, err := zipkinc.NewTracer("product", 10000)
	if err != nil {
		log.Fatalf("failed to create tracer: %v", err)
	}
	defer reporter.Close()

	// Start the root span
	rootSpan := tracer.StartSpan("product")
	rootSpan.Tag("product", "product")

	time.Sleep(time.Second * 1)

	// Start a child span for the DB operation
	dbSpan := tracer.StartSpan("db", zipkin.Parent(rootSpan.Context()))
	dbSpan.Tag("db", "db")

	time.Sleep(time.Second * 2)

	// Finish the child span
	dbSpan.Finish()

	// Start another child span for Kafka
	kafkaSpan := tracer.StartSpan("kafka", zipkin.Parent(rootSpan.Context()))
	kafkaSpan.Tag("kafka", "kafka")

	time.Sleep(time.Second * 1)

	// Finish the Kafka span
	kafkaSpan.Finish()

	// Finish the root span
	rootSpan.Finish()
}
