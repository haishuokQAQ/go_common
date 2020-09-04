package gin_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"math"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	MetricNameHttpLatency = "http_request_latency"
	MetricNameCallerCount = "http_caller_count"
	MetricNameRestResult  = "http_response_rest_code"
)

const (
	AccessUserNameHeader       = "webauth-user"
	CallingStationHeader       = "CALLING-STATION"
	AccessTokenHeader          = "x-access-username"
	LoadBalanceForwardIPHeader = "X-Forwarded-For"
	RestCodeHeader             = "Rest-Code"
	LoadBalanceRealIP          = "X-Real-IP"
)

var uvMap = &sync.Map{}

var httpLatencyVec *prometheus.HistogramVec
var httpResultVec *prometheus.CounterVec

func init() {
	httpLatencyVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem:   "",
		Name:        MetricNameHttpLatency,
		Help:        "Http Request Latency",
		ConstLabels: nil,
		Buckets:     []float64{1, 5, 20, 50, 100, 200, 300, 400, 500, 700, 1000, 2000, 5000},
	}, []string{"path", "method"})
	httpResultVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem:   "",
		Name:        MetricNameRestResult,
		Help:        "Http Response Rest Code",
		ConstLabels: nil,
	}, []string{"path", "rest_code"})
	prometheus.DefaultRegisterer.MustRegister(httpLatencyVec)
	prometheus.DefaultRegisterer.MustRegister(httpResultVec)
}

func Monitoring() gin.HandlerFunc {

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				stackInfo := fmt.Sprintf("%s", buf[:n])
				logrus.Errorf("Panic occurred in monitoring middleware!.Panic stack %+v", string(stackInfo))
			}
		}()
		stop := time.Since(start)
		statusCode := strconv.Itoa(c.Writer.Status())
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		httpLatencyVec.With(map[string]string{
			"path":   path,
			"method": c.Request.Method,
		}).Observe(float64(latency))
			httpResultVec.With(map[string]string{
				"path":      path,
				"rest_code": statusCode,
			}).Inc()
	}
}
