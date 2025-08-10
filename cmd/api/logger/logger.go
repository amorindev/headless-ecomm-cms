package logger

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggerMdw(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
		logrus.SetOutput(os.Stdout)

		latency := time.Since(start)
		statusCode := rw.status
		method := r.Method
		path := r.URL.Path

		var statusColor string
		switch {
		case statusCode >= 200 && statusCode < 300:
			statusColor = "\033[32m" 
		case statusCode >= 300 && statusCode < 400:
			statusColor = "\033[36m" 
		case statusCode >= 400 && statusCode < 500:
			statusColor = "\033[33m" 
		default:
			statusColor = "\033[31m"
		}

		resetColor := "\033[0m"

		logrus.Infof(
			"%s | %s %3d %s | %13v | %-7s %s",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusColor, statusCode, resetColor,
			latency,
			method,
			path,
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
