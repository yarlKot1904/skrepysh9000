package node_exporter

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"skrepysh-agent/pkg/config"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/prom2json"
)

func GetNodeExporterMetricsJSON(config *config.NodeExporterConfig) ([]byte, error) {
	mfChan := make(chan *dto.MetricFamily, 1024)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	u, err := url.JoinPath(fmt.Sprintf("%s:%d", config.Host, config.Port), "/metrics")
	if err != nil {
		return nil, err
	}
	err = prom2json.FetchMetricFamilies(u, mfChan, transport)
	if err != nil {
		return nil, err
	}

	var result = make(map[string]*prom2json.Family)
	for mf := range mfChan {
		switch *mf.Name {
		case "node_memory_MemAvailable_bytes", "node_memory_MemTotal_bytes",
			"node_filesystem_avail_bytes", "node_filesystem_size_bytes",
			"node_cpu_seconds_total", "node_memory_Percpu_bytes":
			result[*mf.Name] = prom2json.NewFamily(mf)
		default:
		}
	}

	jsonText, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return jsonText, nil
}
