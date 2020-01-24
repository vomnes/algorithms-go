.PHONY: build doc fmt lint run test vendor_clean vendor_get

GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

default: run

build:
	@go build -v -o ./bin/graph-algorithms ./src/*.go

exec:
	@./bin/graph-algorithms

run: build exec
