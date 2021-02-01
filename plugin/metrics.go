package plugin

import (
	"fmt"

	datadog "github.com/DataDog/opencensus-go-exporter-datadog"
	"github.com/ipfs/go-ipfs/plugin"
	"go.opencensus.io/stats/view"
)

var _ plugin.Plugin = &MetricsPlugin{}

type MetricsPlugin struct {
	exporter *datadog.Exporter
}

func (m MetricsPlugin) Name() string {
	return "datadog-metrics"
}

func (m MetricsPlugin) Version() string {
	return "0.0.1"
}

func (m *MetricsPlugin) Init(env *plugin.Environment) error {
	dd, err := datadog.NewExporter(datadog.Options{
		Namespace: "go_ipfs",
	})
	if err != nil {
		return fmt.Errorf("failed to create the Datadog exporter: %v", err)
	}

	m.exporter = dd

	view.RegisterExporter(dd)

	return nil
}

func (m *MetricsPlugin) Close() error {
	m.exporter.Stop()
	return nil
}
