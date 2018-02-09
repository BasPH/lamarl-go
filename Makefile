SHELL := /bin/bash
GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GIT_COMMIT=$(shell git rev-parse --short HEAD)
IMAGE_NAME := basph/kraken-exporter

.PHONY: all fmt vet lint build

all: fmt vet lint build

fmt:
	@gofmt -s -l -w $(GOFILES)

vet:
	@go tool vet $(GOFILES)

lint:
	@for f in $(GOFILES); do golint $d; done
