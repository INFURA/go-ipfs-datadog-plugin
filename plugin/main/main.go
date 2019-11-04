package main

import (
	plugin "github.com/INFURA/go-ipfs-datadog-plugin/plugin"
)

var Plugins = plugin.Plugins //nolint

func main() {
	panic("this is a plugin, build it as a plugin, this is here as for go#20312")
}
