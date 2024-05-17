package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"

	"skrepysh-agent/pkg/config"
	"skrepysh-agent/pkg/node_exporter"
)

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Serve(log *zap.Logger, conf *config.Config) error {
	for e, h := range getHandlers(log, conf) {
		http.HandleFunc(e, h)
	}
	log.Info(fmt.Sprintf("started server on port %d", conf.ServerPort))
	return http.ListenAndServe(fmt.Sprintf(":%d", conf.ServerPort), corsHandler(http.DefaultServeMux))
}

func getHandlers(log *zap.Logger, conf *config.Config) map[string]func(w http.ResponseWriter, r *http.Request) {
	handlers := make(map[string]func(w http.ResponseWriter, r *http.Request))

	const configureEndpoint = "/configure"
	handlers[configureEndpoint] = func(w http.ResponseWriter, r *http.Request) {
		log.Info(fmt.Sprintf("%s request to %s", r.Method, configureEndpoint))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		switch r.Method {
		case http.MethodPost:
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, wrapStatus("Error reading request body"), http.StatusBadRequest)
				return
			}
			request := &configureRequest{}
			err = json.Unmarshal(bodyBytes, request)
			if err != nil {
				http.Error(w, wrapStatus("Error unmarshalling request body"), http.StatusBadRequest)
				return
			}
			err = request.exec(log)
			if err != nil {
				http.Error(w, wrapStatus(fmt.Sprintf("Error running commands: %s", err.Error())),
					http.StatusInternalServerError)
				return
			}
			_, err = w.Write([]byte(wrapStatus("OK")))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, wrapStatus("Method not allowed"), http.StatusMethodNotAllowed)
		}
	}

	const healthcheckEndpoint = "/healthcheck"
	handlers[healthcheckEndpoint] = func(w http.ResponseWriter, r *http.Request) {
		log.Info(fmt.Sprintf("%s request to %s", r.Method, healthcheckEndpoint))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		_, err := w.Write([]byte(wrapStatus("OK")))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	const metricsEndpoint = "/metrics"
	handlers[metricsEndpoint] = func(w http.ResponseWriter, r *http.Request) {
		log.Info(fmt.Sprintf("%s request to %s", r.Method, metricsEndpoint))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		switch r.Method {
		case http.MethodGet:
			metrics, err := node_exporter.GetNodeExporterMetricsJSON(&conf.NodeExporter)
			if err != nil {
				log.Error("error fetching node exporter metrics", zap.Error(err))
				http.Error(w, wrapStatus("error fetching node exporter metrics"), http.StatusInternalServerError)
			}
			_, err = w.Write(metrics)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, wrapStatus("Method not allowed"), http.StatusMethodNotAllowed)
		}
	}

	return handlers
}

func wrapStatus(statusText string) string {
	return fmt.Sprintf("{\"status\": \"%s\"}", statusText)
}
