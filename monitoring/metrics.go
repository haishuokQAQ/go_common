package monitoring

import (
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type CustomCollector struct {
	metricsMap       map[string]*prometheus.Metric
	metricsWriteLock *sync.RWMutex
}

func (collector *CustomCollector) Collect(ch chan<- prometheus.Metric) {
	collector.metricsWriteLock.RLock()
	defer collector.metricsWriteLock.RUnlock()
	for k, _ := range collector.metricsMap {
		ch <- *collector.metricsMap[k]
	}
}

func (collector *CustomCollector) Describe(ch chan<- *prometheus.Desc) {
	collector.metricsWriteLock.RLock()
	defer collector.metricsWriteLock.RUnlock()
	for k, _ := range collector.metricsMap {
		ch <- (*collector.metricsMap[k]).Desc()
	}
}

var DuplicateMetric = errors.New("Duplicate metric!")
var MetricNotFound = errors.New("Metric not found!")
var GetErrorType = errors.New("Get error type of metric!")

func (collector *CustomCollector) RegisterMetric(name string, metric prometheus.Metric) error {
	collector.metricsWriteLock.Lock()
	defer collector.metricsWriteLock.Unlock()
	if _, ok := collector.metricsMap[name]; ok {
		return DuplicateMetric
	}
	collector.metricsMap[name] = &metric
	return nil
}

func (collector *CustomCollector) GetMetricByName(name string) *prometheus.Metric {
	collector.metricsWriteLock.RLock()
	defer collector.metricsWriteLock.RUnlock()
	return collector.metricsMap[name]
}

func (collector *CustomCollector) GetGauge(name string) (prometheus.Gauge, error) {
	metric := collector.GetMetricByName(name)
	if metric == nil {
		return nil, MetricNotFound
	}
	if gauge, ok := (*metric).(prometheus.Gauge); !ok {
		return nil, GetErrorType
	} else {
		return gauge, nil
	}
}

func (collector *CustomCollector) GetCounter(name string) (prometheus.Counter, error) {
	metric := collector.GetMetricByName(name)
	if metric == nil {
		return nil, MetricNotFound
	}
	if counter, ok := (*metric).(prometheus.Counter); !ok {
		return nil, GetErrorType
	} else {
		return counter, nil
	}
}

func (collector *CustomCollector) GetHistogram(name string) (prometheus.Histogram, error) {
	metric := collector.GetMetricByName(name)
	if metric == nil {
		return nil, MetricNotFound
	}
	if histogram, ok := (*metric).(prometheus.Histogram); !ok {
		return nil, GetErrorType
	} else {
		return histogram, nil
	}
}

func (collector *CustomCollector) GetSummary(name string) (prometheus.Summary, error) {
	metric := collector.GetMetricByName(name)
	if metric == nil {
		return nil, MetricNotFound
	}
	if summary, ok := (*metric).(prometheus.Summary); !ok {
		return nil, GetErrorType
	} else {
		return summary, nil
	}
}

func (collector *CustomCollector) NewCounter(opts prometheus.CounterOpts) (prometheus.Counter, error) {
	counter := prometheus.NewCounter(opts)
	err := collector.RegisterMetric(opts.Name, prometheus.Metric(counter))
	if err != nil {
		return nil, err
	}
	return counter, nil
}

func (collector *CustomCollector) NewGauge(opts prometheus.GaugeOpts) (prometheus.Gauge, error) {
	gauge := prometheus.NewGauge(opts)
	err := collector.RegisterMetric(opts.Name, prometheus.Metric(gauge))
	if err != nil {
		return nil, err
	}
	return gauge, nil
}

func (collector *CustomCollector) NewHistogram(opts prometheus.HistogramOpts) (prometheus.Histogram, error) {
	histogram := prometheus.NewHistogram(opts)
	err := collector.RegisterMetric(opts.Name, prometheus.Metric(histogram))
	if err != nil {
		return nil, err
	}
	return histogram, nil
}

func (collector *CustomCollector) NewSummary(opts prometheus.SummaryOpts) (prometheus.Summary, error) {
	summary := prometheus.NewSummary(opts)
	err := collector.RegisterMetric(opts.Name, prometheus.Metric(summary))
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func NewCustomCollector() *CustomCollector {
	return &CustomCollector{
		metricsMap:       map[string]*prometheus.Metric{},
		metricsWriteLock: &sync.RWMutex{},
	}
}
