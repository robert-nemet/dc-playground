package middleware

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
	}, []string{"status_code", "path"})
)

func init() {
	prometheus.MustRegister(histogram)
}

type wrapresponsewriter struct {
	http.ResponseWriter
	statusCode int
}

func Instrument(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrapresponsewriter{w, http.StatusOK}
		handler(wrapped, r)
		histogram.WithLabelValues(http.StatusText(wrapped.statusCode), r.URL.Path).Observe(time.Since(start).Seconds())
	}
}

func (w *wrapresponsewriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
