package plugin

import "github.com/ipfs/kubo/plugin"

var Plugins = []plugin.Plugin{
	&LoggerPlugin{},
	&MetricsPlugin{},
}
