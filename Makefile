.DEFAULT_GOAL := help
SHELL := /bin/bash

#help: @ list available tasks on this project
help:
	@grep -E '[a-zA-Z\.\-]+:.*?@ .*$$' $(MAKEFILE_LIST)| tr -d '#'  | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#doc: @ generate command documentation
doc:
	@go run cmd/doc/main.go

#test.unit: @ run unit tests and coverage
test.unit:
	@echo "[TEST.UNIT] run unit tests and coverage"
	@go test -race -covermode=atomic -coverprofile=coverage.out \
		github.com/bigblueswarm/bbsctl/pkg/admin \
		github.com/bigblueswarm/bbsctl/pkg/cmd/apply \
		github.com/bigblueswarm/bbsctl/pkg/cmd/clusterinfo \
		github.com/bigblueswarm/bbsctl/pkg/cmd/delete \
		github.com/bigblueswarm/bbsctl/pkg/cmd/describe \
		github.com/bigblueswarm/bbsctl/pkg/cmd/get \
		github.com/bigblueswarm/bbsctl/pkg/cmd/init \
		github.com/bigblueswarm/bbsctl/pkg/cmd/root \
		github.com/bigblueswarm/bbsctl/pkg/config \
		github.com/bigblueswarm/bbsctl/pkg/render\
		github.com/bigblueswarm/bbsctl/pkg/system

#build: @ build bbsctl binary
build:
	@echo "[BUILD] build bbsctl binary"
	rm -rf bin
	go build -o ./bin/bbsctl ./cmd/bbsctl/main.go