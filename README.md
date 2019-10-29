# go-ipfs-datadog-plugin

Datadog monitoring plugin for go-ipfs.


## Usage

```
make install
```


## Caveats

- Make sure go-ipfs is compiled using the same Go and module versions as this
  plugin.
- If you are using go-ipfs 0.4.22 or older, some traces will be lost. See:
  https://github.com/ipfs/go-ipfs/pull/6672


## References

- Boilerplate for this repo is based on https://github.com/ipfs/go-ipfs-example-plugin


## License

MIT
