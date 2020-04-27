mkfile := $(abspath $(lastword $(MAKEFILE_LIST)))
dir := $(dir $(mkfile))

export LOG_TRUNCATE=true
export LOG_DIRECTORY=$(dir)

.PHONY: test
test:
	@go test --cover github.com/digitalocean/concourse-resource-library/...
