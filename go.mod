module github.com/INFURA/go-datadog-plugin

go 1.12

require (
	github.com/DataDog/datadog-go v2.2.0+incompatible // indirect
	github.com/ipfs/go-ipfs v0.4.22
	github.com/opentracing/opentracing-go v1.1.0
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/tinylib/msgp v1.1.0 // indirect
	golang.org/x/sys v0.0.0-20190922100055-0a153f010e69 // indirect
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.18.0
)

// fork to support proper closing of Tracer plugins
// https://github.com/ipfs/go-ipfs/pull/6672
replace github.com/ipfs/go-ipfs => github.com/INFURA/go-ipfs v0.0.0-20190924053227-e40b24f78766
