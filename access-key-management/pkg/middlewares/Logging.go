package middlewares

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getAccessKey(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("missing or invalid Authorization header")
	}

	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

var fileLogger *log.Logger

func InitLogger() error {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	fileLogger = log.New(logFile, "", log.LstdFlags)
	return nil
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := getAccessKey(r)

		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		// Process request
		next.ServeHTTP(lrw, r)

		duration := time.Since(start)
		fileLogger.Printf("%s %s %d %s %s", r.Method, r.URL.Path, lrw.statusCode, duration, token)
	})
}
