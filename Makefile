mkfile := $(abspath $(lastword $(MAKEFILE_LIST)))
dir := $(dir $(mkfile))

export LOG_TRUNCATE=true
export LOG_DIRECTORY=$(dir)

.PHONY: test
test:
	@go test --cover github.com/digitalocean/concourse-resource-library/...

.PHONY: get-bindata
get-bindata:
	go get -u github.com/go-bindata/go-bindata/...

.PHONY: compile-templates
compile-templates:
	go-bindata -pkg bootstrap -o bootstrap/bindata.go bootstrap/src/
	gofmt -w bootstrap/templates.go
