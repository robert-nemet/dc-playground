package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
	}, []string{"code", "method", "path"})
)

type Middleware interface {
	Instrument(next http.Handler) http.Handler
}

type middleware struct {
}

func NewMiddleware() Middleware {
	prometheus.MustRegister(duration)

	return &middleware{}
}

type wrapresponsewriter struct {
	http.ResponseWriter
	statusCode int
}

func (m *middleware) Instrument(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &wrapresponsewriter{ResponseWriter: w, statusCode: 200}
		start := time.Now()
		next.ServeHTTP(wrapped, r)
		end := time.Now()
		duration.WithLabelValues(strconv.Itoa(wrapped.statusCode), r.Method, r.URL.Path).Observe(end.Sub(start).Seconds())
	})
}

func (w *wrapresponsewriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
