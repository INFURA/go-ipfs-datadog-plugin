package plugin

import (
	"testing"

	"github.com/ipfs/kubo/plugin"
	"github.com/stretchr/testify/require"
)

func TestProfilerPluginDefault(t *testing.T) {
	conf := &profilerConfig{}
	tp := &ProfilerPlugin{}
	err := tp.Init(&plugin.Environment{Config: conf})
	require.NoError(t, err)
	require.Equal(t, defaultTracerName, tp.conf.ProfilerName)
}

func TestProfilerPluginConfig(t *testing.T) {
	conf := &profilerConfig{
		ProfilerName: "foo",
	}
	tp := &ProfilerPlugin{}
	err := tp.Init(&plugin.Environment{Config: conf})
	require.NoError(t, err)
	require.Equal(t, "foo", tp.conf.ProfilerName)
}
