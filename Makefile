version = v1.0.0-rc.1

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)

.PHONY: help

help: ## display help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build  for your current architecture
	go build -o bin/hosts2dnsmasq-$(version)-$(GOOS)-$(GOARCH)

.PHONY: build-linux-mipsle
build-linux-mipsle: ## build for mipsle architecture (like Ubiquiti ER-X)
	GOOS=linux GOARCH=mipsle go build -o bin/hosts2dnsmasq-$(version)-linux-mipsle

.PHONY: build-linux-amd64
build-linux-amd64: ## build for amd64 architecture
	GOOS=linux GOARCH=amd64 go build -o bin/hosts2dnsmasq-$(version)-linux-amd64

.PHONY: build-linux-armv5
build-linux-armv5: ## build for armv5 architecture (like Raspberry Pi)
	GOOS=linux GOARCH=arm GOARM=5 go build -o bin/hosts2dnsmasq-$(version)-linux-armv5

.PHONY: build-cross
build-cross: build-linux-amd64 build-linux-mipsle build-linux-armv5 ## build for all architectures

.PHONY: lint
lint: ## run golint
	$(GOPATH)/bin/golint -set_exit_status ./...

.PHONY: golint
golint: ## get golint tool
	export GO111MODULE=off; go get -u golang.org/x/lint/golint

.PHONY: test
test: ## run all tests with coverage report
	go test -cover ./...

.PHONY: clean
clean: ## deletes binaries generated by build*
	rm -Rf ./bin/*