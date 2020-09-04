package monitoring

import (
	"bytes"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"testing"
)

func TestNewCustomCollector(t *testing.T) {
	reg := prometheus.DefaultRegisterer
	collector := NewCustomCollector()
	gauge, err := collector.NewGauge(prometheus.GaugeOpts{
		Name: "TEST",
		Help: "TEST METRIC",
		ConstLabels: map[string]string{
			"srv": "test",
			"a":   "b",
		},
	})
	if err != nil {
		panic(err)
	}
	if err := reg.Register(collector); err != nil {
		panic(err)
	}

	gauge.Add(13)
	buf := bytes.NewBuffer(make([]byte, 40960))
	enc := expfmt.NewEncoder(buf, expfmt.FmtText)
	mfs, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		panic(err)
	}
	for _, mf := range mfs {
		if err := enc.Encode(mf); err != nil {
			panic(err)
		}
	}
	fmt.Println(buf.String())
}
