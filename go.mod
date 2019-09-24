module github.com/INFURA/go-datadog-plugin

go 1.12

require (
	github.com/ipfs/go-ipfs v0.4.22
	github.com/ipfs/go-log v0.0.1
	github.com/opentracing/opentracing-go v1.1.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.18.0
)

// fork to support proper closing of Tracer plugins
// https://github.com/ipfs/go-ipfs/pull/6672
replace github.com/ipfs/go-ipfs => github.com/INFURA/go-ipfs v0.0.0-20190924082731-f1dc529d6327
