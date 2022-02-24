package plugin

import (
	"encoding/json"

	"github.com/ipfs/go-ipfs/plugin"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

var _ plugin.Plugin = &ProfilerPlugin{}

const defaultProfilerName = "go-ipfs"

type profilerConfig struct {
	ProfilerName string `json:"ProfilerName"`
}

type ProfilerPlugin struct {
	conf profilerConfig
}

func (p *ProfilerPlugin) Name() string {
	return "datadog-profiler"
}

func (p *ProfilerPlugin) Version() string {
	return "0.0.1"
}

func (p *ProfilerPlugin) Init(env *plugin.Environment) error {
	err := p.loadConfig(env)
	if err != nil {
		return err
	}

	return profiler.Start(
		profiler.WithService(p.conf.ProfilerName),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
		),
	)
}

func (p *ProfilerPlugin) Close() error {
	profiler.Stop()
	return nil
}

func (p *ProfilerPlugin) loadConfig(env *plugin.Environment) error {
	// load config data
	bytes, err := json.Marshal(env.Config)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, &p.conf); err != nil {
		return err
	}
	if p.conf.ProfilerName == "" {
		p.conf.ProfilerName = defaultProfilerName
	}
	return nil
}
