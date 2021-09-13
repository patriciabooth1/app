package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/patriciabooth1/app/logging"
	"github.com/patriciabooth1/app/models"
	"net/http"
	"time"
)

type testHandler struct {
	log logging.Logger
}

func newTestHandler(log logging.Logger) *testHandler {
	return &testHandler{
		log: log,
	}
}

func (c *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c.log.InfoRequest("Test request received", r)

	startTime := time.Now()

	var body models.TestRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.WarnRequest("Failed to decode request body", r)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c.log.TraceRequest("Data received", &body, r)

	c.log.InfoRequest("Test request processed successfully", r)

	timeToProcessRequest := time.Now().Sub(startTime)

	c.log.DebugRequest(fmt.Sprintf("Request took %v to process", timeToProcessRequest), r)

	w.WriteHeader(http.StatusOK)
}
