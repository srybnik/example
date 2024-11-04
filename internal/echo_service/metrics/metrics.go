package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	s3FreeSpace prometheus.Gauge
}

func New() *Metrics {
	metrics := &Metrics{
		s3FreeSpace: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "free_space_in_s3",
			Help: "Free space in storage",
		}),
	}

	prometheus.MustRegister(metrics.s3FreeSpace)

	return metrics
}

func (m *Metrics) SetS3FreeSpace(space int64) {
	m.s3FreeSpace.Set(float64(space))
}
