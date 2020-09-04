package gin

import (
	"bytes"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"syscall"
)

func RegisterMetricRouterForGin(ginRouterGroup *gin.RouterGroup, metricsPath string) {
	ginRouterGroup.GET(metricsPath, func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				resultData := "#Error %+v"
				if r == nil {
					resultData = fmt.Sprintf(resultData, " ")
				} else {
					resultData = fmt.Sprintf(resultData, r)
				}
				c.Data(http.StatusOK, "text/plain; version=0.0.4; charset=utf-8", []byte(resultData))
			}
		}()
		gather := prometheus.DefaultGatherer
		buf := bytes.NewBuffer(make([]byte, 0))
		enc := expfmt.NewEncoder(buf, expfmt.FmtText)
		mfs, err := gather.Gather()
		if err != nil {
			panic(err)
		}
		for _, mf := range mfs {
			if err := enc.Encode(mf); err != nil {
				panic(err)
			}
		}
		c.Data(http.StatusOK, "text/plain; version=0.0.4; charset=utf-8", []byte(buf.String()))
	})
}

func StartHealthServer(port int) {
	r := gin.New()
	baseGroup := r.Group("/")
	RegisterMetricRouterForGin(baseGroup, "/debug/monitoring")
	restServer := endless.NewServer(fmt.Sprintf(":%d", port), r)
	restServer.BeforeBegin = func(addr string) {
		log.Printf("health server listening on host: %s. Actual pid is %d", addr, syscall.Getpid())
	}
	restServer.RegisterOnShutdown(func() {
		log.Printf("Server on %d stopped", port)
	})
	g := errgroup.Group{}
	g.Go(func() error {
		return restServer.ListenAndServe()
	})
}
