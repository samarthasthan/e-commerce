package broker

import (
	"log"
	"net/http"
	"sync"
)

const (
	PORT = ":7000"
)

var (
	wg sync.WaitGroup
)

type BrokerServer struct {
	m http.ServeMux
}

func NewBrokerServer() BrokerServer {
	return BrokerServer{}
}

func (b *BrokerServer) Connect() {
	wg.Add(1)

	b.m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Weclome to Broker"))
	})

	go func() {
		if err := http.ListenAndServe(PORT, &b.m); err != nil {
			panic(err)
		}
		wg.Done()
	}()

	log.Printf("Broker listing at %v", PORT)
	wg.Wait()
}
