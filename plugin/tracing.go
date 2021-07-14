package plugin

import (
	"encoding/json"

	"github.com/ipfs/go-ipfs/plugin"
	logging "github.com/ipfs/go-log"
	"github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var _ plugin.PluginTracer = &TracingPlugin{}

var tracerName = "go-ipfs"

type datadogConfig struct {
	TracerName string
}

type TracingPlugin struct{}

func (d *TracingPlugin) Name() string {
	return "datadog-tracer"
}

func (d *TracingPlugin) Version() string {
	return "0.0.1"
}

func (d *TracingPlugin) Init(env *plugin.Environment) error {
	if env == nil || env.Config == nil {
		return nil
	}
	bytes, err := json.Marshal(env.Config)
	if err != nil {
		return err
	}
	config := datadogConfig{}
	if err := json.Unmarshal(bytes, &config); err != nil {
		return err
	}
	if config.TracerName != "" {
		tracerName = config.TracerName
	}
	return nil
}

func (d *TracingPlugin) InitTracer() (opentracing.Tracer, error) {
	return opentracer.New(
		tracer.WithServiceName(tracerName),
		tracer.WithLogger(logger{}),
		tracer.WithRuntimeMetrics(),
		tracer.WithAnalytics(true),
		tracer.WithGlobalTag(ext.ResourceName, "default"),
	), nil
}

func (d *TracingPlugin) Close() error {
	tracer.Stop()
	return nil
}

var log = logging.Logger("datadog")

// logger is an adaptor between ddtrace.Logger and go-log
type logger struct{}

func (logger) Log(msg string) {
	log.Info(msg)
}
