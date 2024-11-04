package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	ioprometheusclient "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	prometheus.DefaultRegisterer = prometheus.NewRegistry()

	metrics := New()

	metrics.SetS3FreeSpace(int64(10))

	metric := ioprometheusclient.Metric{}
	err := metrics.s3FreeSpace.Write(&metric)
	assert.NoError(t, err)
	assert.Equal(t, float64(10), metric.Gauge.GetValue())
}
