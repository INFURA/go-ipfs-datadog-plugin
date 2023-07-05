package plugintesting

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	envVarNameOTELMetricsExporter = "OTEL_METRICS_EXPORTER"
	envVarNameOTELTracesExporter  = "OTEL_TRACES_EXPORTER"
	envVarNameOTLPEndpoint        = "OTEL_EXPORTER_OTLP_ENDPOINT"
	envVarNameOTLPProtocol        = "OTEL_EXPORTER_OTLP_PROTOCOL"
	envVarNameOTLPMetricsProtocol = "OTEL_EXPORTER_OTLP_METRICS_PROTOCOL"

	defaultOTELMetricsExporter = "otlp"
	defaultOTELTracesExporter  = "otlp"
	defaultOTLPEndpoint        = "http://localhost:4317"
	defaultOTLPProtocol        = "grpc"
)

type TestEnv struct{}

type EnvOption func(*testing.T, *TestEnv)

var (
	_ EnvOption = WithOTELMetricsExporter("")
	_ EnvOption = WithOTELTracesExporter("")
	_ EnvOption = WithOTLPEndpoint("")
	_ EnvOption = WithOTLPProtocol("")
	_ EnvOption = WithOTLPMetricsProtocol("")
)

func WithOTELMetricsExporter(v string) EnvOption {
	return newTestEnvVar(envVarNameOTELMetricsExporter, v)
}

func WithOTELTracesExporter(v string) EnvOption {
	return newTestEnvVar(envVarNameOTELMetricsExporter, v)
}

func WithOTLPEndpoint(v string) EnvOption {
	return newTestEnvVar(envVarNameOTLPEndpoint, v)
}

func WithOTLPProtocol(v string) EnvOption {
	return newTestEnvVar(envVarNameOTLPProtocol, v)
}

func WithOTLPMetricsProtocol(v string) EnvOption {
	return newTestEnvVar(envVarNameOTLPMetricsProtocol, v)
}

func NewTestEnv(t *testing.T, opts ...EnvOption) *TestEnv {
	t.Helper()

	testEnv := &TestEnv{}

	for _, k := range []string{
		envVarNameOTELMetricsExporter,
		envVarNameOTELTracesExporter,
		envVarNameOTLPEndpoint,
		envVarNameOTLPProtocol,
		envVarNameOTLPMetricsProtocol,
	} {
		t.Setenv(k, "")
	}

	for _, opt := range opts {
		opt(t, testEnv)
	}

	return testEnv
}

func NewDefaultTestEnv(t *testing.T) *TestEnv {
	t.Helper()

	return NewTestEnv(
		t,
		WithOTELMetricsExporter(defaultOTELMetricsExporter),
		WithOTLPEndpoint(defaultOTLPEndpoint),
		WithOTLPProtocol(defaultOTLPProtocol),
	)
}

func NewTestEnvFromOS(t *testing.T) *TestEnv {
	t.Helper()

	requireEnvVar(t, envVarNameOTELMetricsExporter)
	requireEnvVar(t, envVarNameOTLPEndpoint)
	requireEnvVar(t, envVarNameOTLPProtocol)

	return &TestEnv{}
}

func newTestEnvVar(k, v string) EnvOption {
	return func(t *testing.T, _ *TestEnv) {
		t.Helper()

		t.Setenv(k, v)
	}
}

func requireEnvVar(t *testing.T, k string) {
	t.Helper()

	_, ok := os.LookupEnv(k)
	require.True(t, ok)
}
