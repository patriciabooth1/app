package logging

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	requestIDHeader = "X-request-id"
	contextField    = "context"
	requestIDField  = "request_id"
	statusCodeField = "status_code"
	dataField       = "data"
	errorField      = "error"
)

// SanitizedData provides an interface by which to pass sanitized log data
type SanitizedData interface {
	ToSanitizedJSONString() string
}

// Logger provides an interface by which to log at different levels
type Logger interface {
	logStartOfRequest(req *http.Request)
	logEndOfRequest(req *http.Request, statusCode int)
	ErrorRequest(msg string, err error, req *http.Request)
	Error(msg string, err error)
	WarnRequest(msg string, req *http.Request)
	InfoRequest(msg string, req *http.Request)
	Info(msg string)
	DebugRequest(msg string, req *http.Request)
	TraceRequest(msg string, data SanitizedData, req *http.Request)
}

// NewLogger returns a new concrete implementation of the Logger interface
func NewLogger(level string) (Logger, error) {

	logLevel := log.InfoLevel

	logger := log.New()
	logger.SetLevel(logLevel)

	if level != "" {
		logLevel, err := log.ParseLevel(strings.ToLower(level))
		if err != nil {
			return nil, err
		}
		logger.SetLevel(logLevel)
	}

	return &LoggerImpl{
		logger: logger,
	}, nil
}

// LoggerImpl implements the Logger interface
type LoggerImpl struct {
	logger *log.Logger
}

func (i *LoggerImpl) logStartOfRequest(req *http.Request) {

	requestID := uuid.New().String()

	req.Header.Set(requestIDHeader, requestID)

	i.logger.WithFields(log.Fields{
		contextField:   req.Method + " " + req.RequestURI,
		requestIDField: requestID,
	}).Info("Start of request")
}

func (i *LoggerImpl) logEndOfRequest(req *http.Request, statusCode int) {

	i.logger.WithFields(log.Fields{
		contextField:    req.Method + " " + req.RequestURI,
		requestIDField:  req.Header.Get(requestIDHeader),
		statusCodeField: statusCode,
	}).Info("End of request")
}

// ErrorRequest logs an error with an accompanying message, along with request context
func (i *LoggerImpl) ErrorRequest(msg string, err error, req *http.Request) {

	i.logger.WithFields(log.Fields{
		contextField:   req.Method + " " + req.RequestURI,
		requestIDField: req.Header.Get(requestIDHeader),
		errorField:     err.Error(),
	}).Error(msg)
}

// Error logs an error with an accompanying message
func (i *LoggerImpl) Error(msg string, err error) {

	i.logger.WithFields(log.Fields{
		errorField: err.Error(),
	}).Error(msg)
}

// WarnRequest logs a message at warn level, along with request context
func (i *LoggerImpl) WarnRequest(msg string, req *http.Request) {

	i.logger.WithFields(log.Fields{
		contextField:   req.Method + " " + req.RequestURI,
		requestIDField: req.Header.Get(requestIDHeader),
	}).Warn(msg)
}

// InfoRequest logs a message at info level, along with request context
func (i *LoggerImpl) InfoRequest(msg string, req *http.Request) {

	i.logger.WithFields(log.Fields{
		contextField:   req.Method + " " + req.RequestURI,
		requestIDField: req.Header.Get(requestIDHeader),
	}).Info(msg)
}

// Info logs a message at info level
func (i *LoggerImpl) Info(msg string) {

	i.logger.Info(msg)
}

// DebugRequest logs a message at debug level, along with request context
func (i *LoggerImpl) DebugRequest(msg string, req *http.Request) {

	i.logger.WithFields(log.Fields{
		contextField:   req.Method + " " + req.RequestURI,
		requestIDField: req.Header.Get(requestIDHeader),
	}).Debug(msg)
}

// TraceRequest logs a message at trace level, along with request context and a sanitized data object
func (i *LoggerImpl) TraceRequest(msg string, data SanitizedData, req *http.Request) {

	i.logger.WithFields(log.Fields{
		contextField:   req.Method + " " + req.RequestURI,
		requestIDField: req.Header.Get(requestIDHeader),
		dataField:      data.ToSanitizedJSONString(),
	}).Trace(msg)
}
