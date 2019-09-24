# Force Go Modules
GO111MODULE = on

GOCC ?= go
GOFLAGS ?=

# make reproducable
GOFLAGS += -asmflags=all=-trimpath="$(GOPATH)" -gcflags=all=-trimpath="$(GOPATH)"

# If set, override the install location for plugins
IPFS_PATH ?= $(HOME)/.ipfs

# If set, override the IPFS version to build against. This _modifies_ the local
# go.mod/go.sum files and permanently sets this version.
IPFS_VERSION ?= $(lastword $(shell $(GOCC) list -m github.com/ipfs/go-ipfs))

.PHONY: install build

# We currently use a forked go-ipfs but
# the script to make sure the plugins have the exact same version as go-ipfs
# actually overide this fork with the normal protocol labs one. We have to disable
# it for now and deal with the version manually.
go.mod: FORCE
#	./set-target.sh $(IPFS_VERSION)

FORCE:

datadog-plugin.so: plugin/main/main.go go.mod
	$(GOCC) build $(GOFLAGS) -buildmode=plugin -o "$@" "$<"
	chmod +x "$@"

build: datadog-plugin.so
	@echo "Built against" $(IPFS_VERSION)

install: build
	install -Dm700 datadog-plugin.so "$(IPFS_PATH)/plugins/datadog.so"
