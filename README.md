# go-ipfs-datadog-plugin

This repository contains the following `go-ipfs` plugins:
- Datadog logger plugin allows users to set log levels for each `go-ipfs` subsystem. 
- OpenTelemetry metrics plugin configures an OTLP exporter that sends metrics to an OpenTelemetry collector.

## Caveats

- Plugins only work on Linux and MacOS at the moment. You can track the progress of this issue here: https://github.com/golang/go/issues/19282

- If you are using go-ipfs 0.4.22 or older, some traces will be lost. See: https://github.com/ipfs/kubo/pull/6672


## Building and Installing

You must build the plugin with the *exact* version of go used to build the go-ipfs binary you will use it with. You can find the go version for go-ipfs builds from dist.ipfs.io in the build-info file, e.g. https://dist.ipfs.io/go-ipfs/v0.4.22/build-info or by running `ipfs version --all`

You can build this plugin by running `make build`. You can then install it into your local IPFS repo by running `make install`.

Plugins need to be built against the correct version of go-ipfs. This package generally tracks the latest go-ipfs release but if you need to build against a different version, please set the `IPFS_VERSION` environment variable.

You can set `IPFS_VERSION` to:

* `vX.Y.Z` to build against that version of IPFS.
* `$commit` or `$branch` to build against a specific go-ipfs commit or branch.
* `/absolute/path/to/source` to build against a specific go-ipfs checkout.

To update the go-ipfs, run:

```bash
> make go.mod IPFS_VERSION=version
```

## Manual Installation

Copy `datadog-plugin.so` to `$IPFS_DIR/plugins/datadog-plugin.so` (or run `make install` if you are installing locally)

### Configuration

Define plugin configurations variables in the ipfs config file.

- datadog-logger config:
```
{
...
"Plugins": {
    "Plugins": {
    ...
      "datadog-logger": {
        "Config": {
            "Levels": {
                "fatal": ["system1", "system2", ...],
                "error": [...]
                "warn": [...]
                ...
            },
            "DefaultLevel": "info"
        },
        "Disabled": false
      },
    ...
    }
  },
...
}
```

- otel-metrics

Like Kubo's `OpenTelemetry`-based tracing, `OpenTelemetry` metrics are configured via environment variables. A Sample `.envrc-sample` file is provided. Make a copy named `.envrc` and follow the instructions to configure the plugin. For other execution environments, these environment variables should be provided via the particular systems environment mechanism (e.g. through a Kubernetes `ConfigMap`.)

## Integration testing

Rudimentary integration testing is provided for the OpenTelemetry-based metrics plugin. To run these tests, a local configuration must be provided via environment variables to connect the OTEL exporter to a working OTEL collector. Run the tests using the following command:

```
make integration
```

The integration tests will take about 5 minutes and should result in the following metrics being sent to your collector:

- `go-ipfs-datadog-plugin.integration_test.counter`
- `go-ipfs-datadog-plugin.integration_test.histogram`

Manually verify the presence of these metrics in your collector.

## References

- Boilerplate for this repo is based on https://github.com/ipfs/go-ipfs-example-plugin

## License

MIT
