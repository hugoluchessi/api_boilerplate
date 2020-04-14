package config

import (
	"os"

	log "github.com/hugoluchessi/gobservable/logging"
	m "github.com/hugoluchessi/gobservable/metrics"
	prom "github.com/hugoluchessi/gobservable/metrics/providers/prometheus"
)

type ServerTools struct {
	Logger        *log.ContextLogger
	MetricService *m.MetricService
}

func NewServerTools() (*ServerTools, error) {
	l := initLogger()
	ms, err := initPrometheusMetrics()

	if err != nil {
		return nil, err
	}

	return &ServerTools{l, ms}, nil
}

func initLogger() *log.ContextLogger {
	cfg := log.LoggerConfig{Output: os.Stdout}
	cfgs := []log.LoggerConfig{cfg}

	zap := log.NewZapLogger(cfgs)
	l := log.NewContextLogger(zap)

	return l
}

func initPrometheusMetrics() (*m.MetricService, error) {
	cfg := m.DefaultConfig("gobservable_test")
	promSink, err := prom.NewSink()

	if err != nil {
		return nil, err
	}

	ms := m.NewMetricService(cfg, promSink)

	if err != nil {
		return nil, err
	}

	return ms, nil
}
