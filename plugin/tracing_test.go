package plugin

import (
	"testing"

	"github.com/ipfs/go-ipfs/plugin"
	"github.com/stretchr/testify/require"
)

func TestTracingPluginDefault(t *testing.T) {
	conf := &tracerConfig{}
	tp := &TracingPlugin{}
	err := tp.Init(&plugin.Environment{Config: conf})
	require.NoError(t, err)
	require.Equal(t, defaultTracerName, tp.conf.TracerName)
}

func TestTracingPluginConfig(t *testing.T) {
	conf := &tracerConfig{
		TracerName: "foo",
	}
	tp := &TracingPlugin{}
	err := tp.Init(&plugin.Environment{Config: conf})
	require.NoError(t, err)
	require.Equal(t, "foo", tp.conf.TracerName)
}
