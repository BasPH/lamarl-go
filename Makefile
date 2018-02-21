SHELL := /bin/bash
GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GIT_COMMIT=$(shell git rev-parse --short HEAD)

.PHONY: all fmt vet lint test build

all: fmt vet lint build

fmt:
	@gofmt -s -l -w $(GOFILES)

vet:
	@go tool vet $(GOFILES)

lint:
	@for f in $(GOFILES); do golint $d; done

build:
	@GOOS=linux go build -o main

test:
	@go test