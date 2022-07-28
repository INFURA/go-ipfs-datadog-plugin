package plugin

import (
	"encoding/json"

	logging "github.com/ipfs/go-log"
	"github.com/ipfs/kubo/plugin"
	"github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var _ plugin.PluginTracer = &TracingPlugin{}

const defaultTracerName = "go-ipfs"

type tracerConfig struct {
	TracerName string `json:"TracerName"`
}

type TracingPlugin struct {
	conf tracerConfig
}

func (t *TracingPlugin) Name() string {
	return "datadog-tracer"
}

func (t *TracingPlugin) Version() string {
	return "0.0.1"
}

func (t *TracingPlugin) Init(env *plugin.Environment) error {
	return t.loadConfig(env)
}

func (t *TracingPlugin) InitTracer() (opentracing.Tracer, error) {
	return opentracer.New(
		tracer.WithServiceName(t.conf.TracerName),
		tracer.WithLogger(logger{}),
		tracer.WithRuntimeMetrics(),
		tracer.WithAnalytics(true),
	), nil
}

func (t *TracingPlugin) Close() error {
	tracer.Stop()
	return nil
}

func (t *TracingPlugin) loadConfig(env *plugin.Environment) error {
	// load config data
	bytes, err := json.Marshal(env.Config)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, &t.conf); err != nil {
		return err
	}
	if t.conf.TracerName == "" {
		t.conf.TracerName = defaultTracerName
	}
	return nil
}

var log = logging.Logger("datadog")

// logger is an adaptor between ddtrace.Logger and go-log
type logger struct{}

func (logger) Log(msg string) {
	log.Info(msg)
}
