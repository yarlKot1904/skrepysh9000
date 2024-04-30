package node_exporter

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"os"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/prom2json"
)

const (
	NODE_EXPORTER_ENDPOINT = "http://localhost:9100/metrics" //HOST_NAME
)

func main() {
	var err error

	mfChan := make(chan *dto.MetricFamily, 1024)

	transport, err := makeTransport()
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		err := prom2json.FetchMetricFamilies(NODE_EXPORTER_ENDPOINT, mfChan, transport)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	result := []*prom2json.Family{}
	for mf := range mfChan {
		result = append(result, prom2json.NewFamily(mf))
	}

	jsonText, err := json.Marshal(result)
	if err != nil {
		log.Fatalln("error marshaling JSON:", err)
	}
	if _, err := os.Stdout.Write(jsonText); err != nil {
		log.Fatalln("error writing to stdout:", err)
	}
	fmt.Println(jsonText)
}

func makeTransport() (*http.Transport, error) {
	var transport *http.Transport
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return transport, nil
}
