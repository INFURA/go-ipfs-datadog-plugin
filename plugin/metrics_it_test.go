package plugin

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/INFURA/go-ipfs-datadog-plugin/plugin/plugintesting"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	metricapi "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

func TestOpenTelemetryIntegration(t *testing.T) {
	plugintesting.Integration(t)

	_ = plugintesting.NewTestEnvFromOS(t)

	plugin := &MetricsPlugin{}
	require.NoError(t, plugin.Init(nil))
	require.IsType(t, metric.NewMeterProvider(), plugin.shutDownMeterProvider)
	t.Cleanup(func() {
		require.NoError(t, plugin.Close())
	})

	meter := otel.Meter("github.com/INFURA/go-ipfs-datadog-plugin")

	counter, err := meter.Int64Counter("go-ipfs-datadog-plugin.integration_test.counter", metricapi.WithDescription("this is a counter"), metricapi.WithUnit("widgets"))
	require.NoError(t, err)

	histogram, err := meter.Int64Histogram("go-ipfs-datadog-plugin.integration_test.histogram", metricapi.WithDescription("this is a histogram"), metricapi.WithUnit("things"))
	require.NoError(t, err)

	attributes := attribute.NewSet(
		attribute.String("env", "local"),
	)

	for i := 0; i < 5; i++ {
		for _, v := range []int64{1, 2, 3, 4, 5, 4, 3, 2, 1} {
			histogram.Record(context.Background(), v, metricapi.WithAttributeSet(attributes))
			time.Sleep(10 * time.Second)
		}

		for _, v := range []int64{1, 2, 3, 4, 5, 4, 3, 2, 1} {
			counter.Add(context.Background(), v, metricapi.WithAttributeSet(attributes))
			time.Sleep(time.Duration(rand.Int63n(100)*50) * time.Millisecond)
		}
	}
}
