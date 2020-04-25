mkfile := $(abspath $(lastword $(MAKEFILE_LIST)))
dir := $(dir $(mkfile))

.PHONY: test
test:
	@go test -v ./...
