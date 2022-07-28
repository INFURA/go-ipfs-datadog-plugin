package plugin

import (
	"testing"

	"github.com/ipfs/kubo/plugin"
	"github.com/stretchr/testify/require"
)

func TestLoggerPluginDefault(t *testing.T) {
	conf := &loggerConfig{}
	lp := &LoggerPlugin{}
	err := lp.Init(&plugin.Environment{Config: conf})
	require.NoError(t, err)
	require.Equal(t, defaultLoggerLevel, lp.conf.DefaultLevel)
}

func TestLoggerPluginConfig(t *testing.T) {
	conf := &loggerConfig{
		DefaultLevel: "info",
	}
	lp := &LoggerPlugin{}
	err := lp.Init(&plugin.Environment{Config: conf})
	require.NoError(t, err)
	require.Equal(t, "info", lp.conf.DefaultLevel)
}
