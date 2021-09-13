package handlers

import (
	"github.com/gorilla/mux"
	"github.com/patriciabooth1/app/logging"
	"net/http"
)

// Register registers handler functions against all available routes
func Register(router *mux.Router, log logging.Logger) {

	loggingInterceptor := logging.NewLoggingInterceptor(log)

	router.HandleFunc("/health-check", healthCheck).Name("health-check")

	testRouter := router.PathPrefix("/test").Subrouter()
	testRouter.Handle("", newTestHandler(log)).Methods(http.MethodPost).Name("test")
	testRouter.Use(loggingInterceptor.LogIntercept)
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
