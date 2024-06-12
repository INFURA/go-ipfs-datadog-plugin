package plugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ipfs/kubo/plugin"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	metricapi "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	envVarMetricsExporter     = "OTEL_METRICS_EXPORTER"
	envVarOTLPProtocol        = "OTEL_EXPORTER_OTLP_PROTOCOL"
	envVarOTLPMetricsProtocol = "OTEL_EXPORTER_OTLP_METRICS_PROTOCOL"
)

const (
	otlpProtocolNameGRPC = "grpc"
	otlpProtocolNameHTTP = "http/protobuf"
)

const (
	metricPluginName     = "otel-metrics"
	metricsPluginVersion = "v0.8.1"
)

type shutDownMeterProvider interface {
	metricapi.MeterProvider
	Shutdown(ctx context.Context) error
}

var _ shutDownMeterProvider = (*noopShutdownMeterProvider)(nil)

type noopShutdownMeterProvider struct {
	metricapi.MeterProvider
}

func (n *noopShutdownMeterProvider) Shutdown(_ context.Context) error {
	return nil
}

var _ plugin.Plugin = (*MetricsPlugin)(nil)

// MetricsPlugin configures the OpenTelemetry metrics exporter into the
// default MeterProvider which can be retrieved via otel.GetMeterProvider.
//
// See: https://pkg.go.dev/github.com/ipfs/kubo/plugin#Plugin
type MetricsPlugin struct {
	shutDownMeterProvider
	exporters []metric.Exporter
}

// Name returns the name of the plugin.
func (m MetricsPlugin) Name() string {
	return metricPluginName
}

// Version returns the version of the plugin.
func (m MetricsPlugin) Version() string {
	return metricsPluginVersion
}

// Init initializes the kubo plugin (from environment variables).
func (m *MetricsPlugin) Init(_ *plugin.Environment) error {
	mp, exps, err := newMeterProvider(context.Background())
	if err != nil {
		return err
	}

	m.shutDownMeterProvider = mp
	m.exporters = exps
	otel.SetMeterProvider(mp)
	otel.SetTextMapPropagator(autoprop.NewTextMapPropagator())

	return nil
}

// Close safely shuts down the plugin.
func (m *MetricsPlugin) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.Shutdown(ctx)
}

func newMeterProvider(ctx context.Context) (shutDownMeterProvider, []metric.Exporter, error) {
	exporters, err := newMetricsExporters(ctx)
	if err != nil {
		return nil, nil, err
	}

	if len(exporters) < 1 {
		return &noopShutdownMeterProvider{noop.NewMeterProvider()}, nil, nil
	}

	var options []metric.Option
	for _, exporter := range exporters {
		options = append(options, metric.WithReader(metric.NewPeriodicReader(exporter)))
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			// This name is uppercase and not prefixed with "ipfs-" to match the
			// tracing service name already included in the external Kubo project.
			//
			// See: https://github.com/ipfs/kubo/blob/a07852a3f0294974b802923fb136885ad077384e/tracing/tracing.go#L66
			semconv.ServiceNameKey.String("Kubo"),
			semconv.TelemetrySDKLanguageGo,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	options = append(options, metric.WithResource(res))

	return metric.NewMeterProvider(options...), exporters, nil
}

var (
	errUnsupportedExporter      = errors.New("unknown or unsupported exporter")
	errUnsupportedProtocol      = errors.New("unknown or unsupported OTLP protocol")
	errBuildingOTLPgRPCExporter = errors.New("building OTLP gRPC exporter")
	errBuildingOTLPHTTPExporter = errors.New("building OTLP HTTP exporter")
	errBuildingStdoutExporter   = errors.New("building STDOUT exporter")
)

// newMetricExporters returns zero or more metrics exporters configured
// strictly though environment variables.  Like kubo's OpenTelemetry
// trace configuration.  See the README for set-up instructions.
func newMetricsExporters(ctx context.Context) ([]metric.Exporter, error) {
	var exporters []metric.Exporter

	exporterEnvVar := os.Getenv(envVarMetricsExporter)
	if exporterEnvVar == "" {
		return exporters, nil
	}

	for _, s := range strings.Split(exporterEnvVar, ",") {

		switch s {
		case "otlp":
			protocol := otlpProtocolNameHTTP
			if v := os.Getenv(envVarOTLPProtocol); v != "" {
				protocol = v
			}
			if v := os.Getenv(envVarOTLPMetricsProtocol); v != "" {
				protocol = v
			}

			switch protocol {
			case otlpProtocolNameHTTP:
				exporter, err := otlpmetrichttp.New(ctx)
				if err != nil {
					return nil, fmt.Errorf("%w: %w", errBuildingOTLPHTTPExporter, err)
				}
				exporters = append(exporters, exporter)
			case otlpProtocolNameGRPC:
				exporter, err := otlpmetricgrpc.New(ctx)
				if err != nil {
					return nil, fmt.Errorf("%w: %w", errBuildingOTLPgRPCExporter, err)
				}
				exporters = append(exporters, exporter)
			default:
				return nil, fmt.Errorf("%w: %s", errUnsupportedProtocol, protocol)
			}
		case "stdout":
			exporter, err := stdoutmetric.New()
			if err != nil {
				return nil, fmt.Errorf("%w: %w", errBuildingStdoutExporter, err)
			}
			exporters = append(exporters, exporter)
		default:
			return nil, fmt.Errorf("%w: %s", errUnsupportedExporter, s)
		}
	}

	return exporters, nil
}
