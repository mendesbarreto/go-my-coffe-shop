package logger

import (
	"net/http"

	"github.com/felixge/httpsnoop"
	"log/slog"
)

func WithLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		slog.Info("http[%d] -- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}
