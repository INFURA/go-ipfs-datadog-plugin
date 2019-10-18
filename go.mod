module github.com/INFURA/go-ipfs-datadog-plugin

go 1.12

require (
	github.com/DataDog/datadog-go v2.2.0+incompatible // indirect
	github.com/INFURA/go-datadog-plugin v0.1.2-0.20190925061845-4ad5667a3177
	github.com/ipfs/go-ipfs v0.4.22
	github.com/ipfs/go-log v0.0.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/tinylib/msgp v1.1.0 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.18.0
)

// fork to support proper closing of Tracer plugins
// https://github.com/ipfs/go-ipfs/pull/6672
replace github.com/ipfs/go-ipfs => github.com/INFURA/go-ipfs v0.0.0-20190926031411-04fabe5e6b7d
