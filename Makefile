version = v1.0.0-rc.1

GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

build:
	go build -o bin/hosts2dnsmasq-$(version)-$(GOOS)-$(GOARCH)

build-mipsle:
	GOOS=linux GOARCH=mipsle go build -o bin/hosts2dnsmasq-$(version)-linux-mipsle

test:
	go test ./...