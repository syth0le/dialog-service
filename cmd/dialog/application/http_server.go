package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	xservers "github.com/syth0le/gopnik/servers"

	"github.com/syth0le/dialog-service/internal/handler/publicapi"
)

func (a *App) newHTTPServer(env *env) *xservers.HTTPServerWrapper {
	return xservers.NewHTTPServerWrapper(
		a.Logger,
		xservers.WithAdminServer(a.Config.AdminServer),
		xservers.WithPublicServer(a.Config.PublicServer, a.publicMux(env)),
	)
}

func (a *App) publicMux(env *env) *chi.Mux {
	mux := chi.NewMux()

	metrics := a.RegisterMetrics()

	handler := &publicapi.Handler{
		Logger:        a.Logger,
		DialogService: env.dialogService,
	}

	mux.Handle("/metrics", promhttp.HandlerFor(a.Registry, promhttp.HandlerOpts{}))

	mux.Route("/dialog", func(r chi.Router) {
		r.Use(env.authClient.AuthenticationInterceptor)
		r.Use(metricsMiddleware(metrics))

		r.Post("/", handler.CreateDialog) // todo: make group dialogs
		r.Post("/send", handler.CreateMessage)
		r.Get("/{dialogID}/list", handler.GetDialogMessages)
	})

	return mux
}

// var (
// 	requestDurationsHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
// 		Namespace: "myapp",
// 		Name:      "request_duration_seconds",
// 		Help:      "Duration of the request.",
// 		// 4 times larger for apdex score
// 		// Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
// 		// Buckets: prometheus.LinearBuckets(0.1, 5, 5),
// 		Buckets: []float64{0.1, 0.15, 0.2, 0.25, 0.3},
// 	}, []string{"status", "method"})
//
// 	requestDurationsHistogram2 = prometheus.NewHistogram(prometheus.HistogramOpts{
// 		Name:    "request_durations_histogram_secs",
// 		Buckets: prometheus.DefBuckets,
// 		Help:    "Requests Durations, in Seconds",
// 	})
// )

// func MetricsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		lrw := NewLoggingResponseWriter(w)
// 		t := prometheus.NewTimer(requestDurationsHistogram.With(prometheus.Labels{
// 			"method": r.Method,
// 			"status": strconv.Itoa(lrw.statusCode),
// 		}))
// 		defer t.ObserveDuration()
// 		t2 := prometheus.NewTimer(requestDurationsHistogram2)
// 		defer t2.ObserveDuration()
// 		next.ServeHTTP(lrw, r)
// 	})
// }
