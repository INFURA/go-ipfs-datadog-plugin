package plugin

import (
	"os"

	"github.com/ipfs/go-ipfs/plugin"
	logging "github.com/ipfs/go-log"
	"github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var log = logging.Logger("datadog")

var Plugins = []plugin.Plugin{
	&DatadogPlugin{},
	&LoggerPlugin{},
}

var _ plugin.PluginTracer = &DatadogPlugin{}

var tracerName = "go-ipfs"

const tracerEnv = "IPFS_TRACER_NAME"

type DatadogPlugin struct{}

func (d *DatadogPlugin) Name() string {
	return "datadog"
}

func (d *DatadogPlugin) Version() string {
	return "0.0.1"
}

func (d *DatadogPlugin) Init(env *plugin.Environment) error {
	maybeName := os.Getenv(tracerEnv)
	if maybeName != "" {
		tracerName = maybeName
	}
	return nil
}

func (d *DatadogPlugin) InitTracer() (opentracing.Tracer, error) {
	return opentracer.New(
		tracer.WithServiceName(tracerName),
		tracer.WithLogger(logger{}),
		tracer.WithRuntimeMetrics(),
		tracer.WithAnalytics(true),
		tracer.WithGlobalTag(ext.ResourceName, "default"),
	), nil
}

func (d *DatadogPlugin) Close() error {
	tracer.Stop()
	return nil
}

// logger is an adaptor between ddtrace.Logger and go-log
type logger struct{}

func (logger) Log(msg string) {
	log.Info(msg)
}
