package plugin

import (
	"os"

	"github.com/ipfs/go-ipfs/plugin"
	"github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var Plugins = []plugin.Plugin{
	&DatadogPlugin{},
}

var _ plugin.PluginTracer = &DatadogPlugin{}

var tracerName = "#TRACER-NAME-NOT-SET"

const tracerEnv = "IPFS_TRACER_NAME"

type DatadogPlugin struct{}

func (d *DatadogPlugin) Name() string {
	return "datadog"
}

func (d *DatadogPlugin) Version() string {
	return "0.0.1"
}

func (d *DatadogPlugin) Init() error {
	maybeName := os.Getenv(tracerEnv)
	if maybeName != "" {
		tracerName = maybeName
	}
	return nil
}

func (d *DatadogPlugin) InitTracer() (opentracing.Tracer, error) {
	// Start the regular tracer and return it as an opentracing.Tracer interface. You
	// may use the same set of options as you normally would with the Datadog tracer.
	return opentracer.New(
		tracer.WithServiceName(tracerName),
		tracer.WithRuntimeMetrics(),
	), nil
}

func (d *DatadogPlugin) Close() error {
	tracer.Stop()
	return nil
}
