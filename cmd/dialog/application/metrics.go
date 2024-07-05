package application

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	labelMethod = "method"
	labelPath   = "path"
	labelCode   = "code"
)

type Metrics struct {
	requestsReceived *prometheus.CounterVec
	requestDuration  *prometheus.HistogramVec
}

func (a *App) RegisterMetrics() *Metrics {
	reqReceived := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "http",
			Name:      "requests_total",
			Help:      "Total number of requests received.",
		},
		[]string{labelMethod, labelPath, labelCode},
	)
	reqDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "http",
			Name:      "request_duration_seconds",
			Help:      "Duration of a request in seconds.",
		},
		[]string{labelMethod, labelPath, labelCode},
	)

	a.Registry.MustRegister(reqReceived)
	a.Registry.MustRegister(reqDuration)

	// Add go runtime metrics and process collectors.
	a.Registry.MustRegister(
		collectors.NewGoCollector(
			collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
		),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	return &Metrics{
		requestsReceived: reqReceived,
		requestDuration:  reqDuration,
	}
}

func (m *Metrics) Default() http.Handler {
	return promhttp.Handler()
}

func (m *Metrics) IncRequests(method, path string, code int) {
	m.requestsReceived.WithLabelValues(method, path, strconv.Itoa(code)).Inc()
}

func (m *Metrics) ObsDuration(method, path string, code int, duration float64) {
	m.requestDuration.WithLabelValues(method, path, strconv.Itoa(code)).Observe(duration)
}

type StatusCodeResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewStatusCodeResponseWriter(w http.ResponseWriter) *StatusCodeResponseWriter {
	return &StatusCodeResponseWriter{w, http.StatusOK}
}

func (lrw *StatusCodeResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func metricsMiddleware(metrics *Metrics) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			lrw := NewStatusCodeResponseWriter(w)

			h.ServeHTTP(lrw, r)

			duration := time.Since(start)
			fmt.Println(duration)
			metrics.IncRequests(r.Method, r.RequestURI, lrw.StatusCode)
			metrics.ObsDuration(r.Method, r.RequestURI, lrw.StatusCode, duration.Seconds())
		})
	}
}
