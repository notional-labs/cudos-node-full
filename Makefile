VERSION := $(shell echo $(shell git describe --tags))
COMMIT := $(shell git log -1 --format='%H')

BUILDDIR ?= $(CURDIR)/build

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=cudos-node \
	-X github.com/cosmos/cosmos-sdk/version.AppName=cudos-noded \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: export CGO_LDFLAGS=-Wl,-rpath,$$ORIGIN/../
install: go.sum
		@echo "--> Installing cudos-noded"
		@go install -mod=readonly $(BUILD_FLAGS) -tags "ledger" ./cmd/cudos-noded


build: export CGO_LDFLAGS=-Wl,-rpath,$$ORIGIN/../
build: go.sum
		@echo "--> Building cudos-noded"
		@go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ -tags "ledger" ./cmd/cudos-noded
		
go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify
    
proto-gen:
	docker run --rm -v $(CURDIR):/workspace --workdir /workspace ghcr.io/cosmos/proto-builder:0.12.1 sh ./scripts/protocgen.sh

test-all: test-unit test-sim-benchmark test-sim-determinism

test-unit:
	@echo "--> Running unit tests"
	@go test -v -mod=readonly -tags='cgo ledger test_ledger_mock norace' ./...

test-sim-benchmark:
	@echo "--> Running benchmark sim tests"
	@go test -v -mod=readonly -benchmem -run ^BenchmarkFullAppSimulation -bench ^BenchmarkFullAppSimulation ./simapp \
		-Enabled=true -NumBlocks=200 -BlockSize=200 -Commit=true -timeout 24h

test-sim-determinism:
	@echo "--> Running determinism sim tests"
	@go test -v -mod=readonly -run ^TestAppStateDeterminism ./simapp \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -timeout 24h
