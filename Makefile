GOPATH:=$(shell go env GOPATH)

.PHONY: build
## build: builds the service
build:
	go build -a -o metrics-svc pkg/*.go

.PHONY: clean
## clean: cleans the binary
clean:
	go clean

.PHONY: run
## run: starts the service
run:
	go run ./pkg

.PHONY: test
## test: executes test suite and generates a coverage file
test:
	go test -v ./... -race -coverprofile=cover.out -covermode=atomic

.PHONY: covreport
## opens the coverage report in the browser
covreport:
	go tool cover -html=cover.out

.PHONY: docker
## docker: builds a docker image with latest tag
docker:
	docker build . -t metrics:latest

.PHONY: format
## format: runs go fmt
format:
	go fmt ./pkg/...

.PHONY: imports
## imports: runs goimports
imports:
	goimports -w -d ./pkg

.PHONY: lint
## lint: runs golangci-lint
lint:
	golangci-lint run

.PHONY: codequality
## format: runs all code quality checks
codequality: clean format imports lint

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
