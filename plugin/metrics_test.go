package plugin

import (
	"testing"

	"github.com/INFURA/go-ipfs-datadog-plugin/plugin/plugintesting"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
)

func NewTestMetricsPlugin(t *testing.T, opts ...plugintesting.EnvOption) *MetricsPlugin {
	t.Helper()

	_ = plugintesting.NewTestEnv(t, opts...)

	return &MetricsPlugin{}
}

func TestMetrics_Init(t *testing.T) {
	t.Log("unit")

	t.Run("error in Init()", func(t *testing.T) {
		testcases := []struct {
			name string
			opts []plugintesting.EnvOption
			err  error
		}{
			{
				name: "with incorrect exporter",
				opts: []plugintesting.EnvOption{
					plugintesting.WithOTELMetricsExporter("wrong"),
				},
				err: errUnsupportedExporter,
			},
			{
				name: "with incorrect protocol",
				opts: []plugintesting.EnvOption{
					plugintesting.WithOTELMetricsExporter("otlp"),
					plugintesting.WithOTLPProtocol("wrong"),
				},
				err: errUnsupportedProtocol,
			},
			{
				name: "with incorrect metrics protocol",
				opts: []plugintesting.EnvOption{
					plugintesting.WithOTELMetricsExporter("otlp"),
					plugintesting.WithOTLPMetricsProtocol("wrong"),
				},
				err: errUnsupportedProtocol,
			},
		}

		for i := range testcases {
			testcase := testcases[i]

			t.Run(testcase.name, func(t *testing.T) {
				require.ErrorIs(t, NewTestMetricsPlugin(t, testcase.opts...).Init(nil), testcase.err)
			})
		}
	})

	t.Run("creates MetricsPlugin with noop exporter", func(t *testing.T) {
		plugin := NewTestMetricsPlugin(t)
		require.Len(t, plugin.exporters, 0)
		require.NoError(t, plugin.Init(nil))
		require.IsType(t, &noopShutdownMeterProvider{}, otel.GetMeterProvider())
		require.NoError(t, plugin.Close())
	})

	t.Run("successfully creates MetricsPlugin", func(t *testing.T) {
		testcases := []struct {
			name string
			opts []plugintesting.EnvOption
		}{
			{
				name: "with default (HTTP) protocol",
				opts: []plugintesting.EnvOption{
					plugintesting.WithOTELMetricsExporter("otlp"),
				},
			},
			{
				name: "with explicit HTTP protocol",
				opts: []plugintesting.EnvOption{
					plugintesting.WithOTELMetricsExporter("otlp"),
					plugintesting.WithOTLPMetricsProtocol("http/protobuf"),
				},
			},
			{
				name: "with explicit gRPC protocol",
				opts: []plugintesting.EnvOption{
					plugintesting.WithOTELMetricsExporter("otlp"),
					plugintesting.WithOTLPMetricsProtocol("grpc"),
				},
			},
		}

		for i := range testcases {
			testcase := testcases[i]
			t.Run(testcase.name, func(t *testing.T) {
				plugin := NewTestMetricsPlugin(t, testcase.opts...)
				require.NoError(t, plugin.Init(nil))
				require.Len(t, plugin.exporters, 1)
			})
		}
	})
}
