package logging

import (
	"net/http"
)

// LoggingInterceptor provides a mechanism by which to log requests and responses
type LoggingInterceptor struct {
	log Logger
}

// NewLoggingInterceptor returns a new LoggingInterceptor
func NewLoggingInterceptor(log Logger) *LoggingInterceptor {
	return &LoggingInterceptor{
		log: log,
	}
}

// LoggingResponseWriter provides a mechanism by which to derive a response code after a request
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader overrides the default implementation of http.ResponseWriter, providing a way to store status code
func (l *LoggingResponseWriter) WriteHeader(code int) {
	l.statusCode = code
	l.ResponseWriter.WriteHeader(code)
}

// LogIntercept logs details of an inbound request and outbound response
func (i *LoggingInterceptor) LogIntercept(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := LoggingResponseWriter{
			ResponseWriter: w,
			statusCode:     200,
		}

		i.log.logStartOfRequest(r)

		next.ServeHTTP(&lw, r)

		i.log.logEndOfRequest(r, lw.statusCode)
	})
}
