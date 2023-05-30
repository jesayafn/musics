package controller

import (
	"github.com/prometheus/client_golang/prometheus"
)

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var RequestResponseTime = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_response_time",
		Help: "Duration of the request",
	},
	[]string{"path"},
)

// type metrics struct {
// 	httpRequestTotal *prometheus.CounterVec
// }

// func NewMetrics(reg prometheus.Registerer) *metrics {
// 	m := &metrics{
// 		httpRequestTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
// 			Name: "http_request_total",
// 			Help: "Total HTTP requests",
// 		}, []string{"path"}),
// 	}
// 	reg.MustRegister(m.httpRequestTotal)
// 	return m
// }

// func RegisterMetric() (*metrics, *prometheus.Registry) {
// 	reg := prometheus.NewRegistry()
// 	metrics := NewMetrics(reg)
// 	// fmt.Sprintf("%p", metrics.httpRequestTotal)
// 	return metrics, reg
// }

// func PrometheusHandler(m *metrics, r *prometheus.Registry) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		promhttp.HandlerFor(r, promhttp.HandlerOpts{Registry: r}).ServeHTTP(ctx.Writer, ctx.Request)
// 	}
// }
