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
		log.Info(fmt.Sprintf("%s request to %s", r.Method, configureEndpoint))
		switch r.Method {
		case http.MethodPost:
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return
			}
			request := &configureRequest{}
			err = json.Unmarshal(bodyBytes, request)
			if err != nil {
				http.Error(w, "Error unmarshalling request body: %s", http.StatusBadRequest)
				return
			}
			err = request.exec()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp := &errorResponse{
					Error:       err.Error(),
					Description: "error running commands",
				}
				respBytes, err := json.Marshal(resp)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Write(respBytes)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

	return handlers
}
