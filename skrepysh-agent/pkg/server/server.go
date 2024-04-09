package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func Serve(log *zap.Logger, port int16) error {
	for e, h := range getHandlers(log) {
		http.HandleFunc(e, h)
	}
	log.Info(fmt.Sprintf("started server on port %d", port))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getHandlers(log *zap.Logger) map[string]func(w http.ResponseWriter, r *http.Request) {
	handlers := make(map[string]func(w http.ResponseWriter, r *http.Request))

	const configureEndpoint = "/configure"
	handlers["/configure"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		log.Info(fmt.Sprintf("%s request to %s", r.Method, configureEndpoint))
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
				http.Error(w, wrapStatus("Error running commands"), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, wrapStatus("Method not allowed"), http.StatusMethodNotAllowed)
		}
	}

	handlers["/healthcheck"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		_, err := w.Write([]byte(wrapStatus("OK")))
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	return handlers
}

func wrapStatus(statusText string) string {
	return fmt.Sprintf("{\"status\": \"%s\"}", statusText)
}
