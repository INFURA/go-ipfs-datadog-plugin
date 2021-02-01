package plugin

import "github.com/ipfs/go-ipfs/plugin"

var Plugins = []plugin.Plugin{
	&TracingPlugin{},
	&LoggerPlugin{},
	&MetricsPlugin{},
}