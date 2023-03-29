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
	}, []string{"status_code", "method", "path"})

	concurentRequests = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "http_concurrent_requests",
	})
)

type Middleware interface {
	Instrument(next http.Handler) http.Handler
}

type middleware struct {
}

func NewMiddleware() Middleware {
	prometheus.MustRegister(duration, concurentRequests)
	return &middleware{}
}

type wrapresponsewriter struct {
	http.ResponseWriter
	statusCode int
}

func (m *middleware) Instrument(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		concurentRequests.Inc()
		start := time.Now()
		wrapped := &wrapresponsewriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(wrapped, r)
		concurentRequests.Dec()
		duration.WithLabelValues(strconv.Itoa(wrapped.statusCode), r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
	})
}

func (w *wrapresponsewriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
