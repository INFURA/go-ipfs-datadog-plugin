# go-ipfs-datadog-plugin

This repository contains the following `go-ipfs` plugins:
- Datadog tracing plugin configures the Datadog tracer to collect the traces and relay them to the agent. `go-ipfs` tracing instrumentation is partial at the moment but should improve over time.
- Datadog continuous profiler pluging configure the Datadog profiler to capture CPU and memory usage over time.
- Datadog logger plugin allows users to set log levels for each `go-ipfs` subsystem. 
- Datadog metrics plugin configure a datadog exporter for OpenCensus metrics.

## Caveats

- Plugins only work on Linux and MacOS at the moment. You can track the progress of this issue here: https://github.com/golang/go/issues/19282

- If you are using go-ipfs 0.4.22 or older, some traces will be lost. See: https://github.com/ipfs/go-ipfs/pull/6672


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

- datadog-tracer config:
```
{
...
"Plugins": {
    "Plugins": {
      ...
      "datadog-tracer": {
        "Config": {
            "TracerName": "go-ipfs-custom"
        },
        "Disabled": false
      }
      ...
    }
  },
...
}
```

- datadog-profiler config:
 ```
...
"Plugins": {
    "Plugins": {
      ...
      "datadog-profiler": {
        "Config": {
            "TracerName": "go-ipfs-custom"
        },
        "Disabled": false
      }
      ...
    }
  },
...
```

## References

- Boilerplate for this repo is based on https://github.com/ipfs/go-ipfs-example-plugin

## License

MIT
