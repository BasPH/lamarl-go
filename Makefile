GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GIT_COMMIT=$(shell git rev-parse --short HEAD)

.PHONY: all fmt vet lint test build update-lambda

all: fmt vet lint build update-lambda

fmt:
	gofmt -s -l -w $(GOFILES)

vet:
	go tool vet $(GOFILES)

lint:
	for f in $(GOFILES); do golint $f; done

build:
	GOOS=linux go build -o main cmd/lambda/*
	go build cmd/sushigo/*

test:
	go test

update-lambda:
	zip main.zip ./main
	aws lambda update-function-code \
    	--function-name awscodestar-lamarl-go-lambda-GetHelloWorld-1KA3SG4UKTUCO \
    	--zip-file fileb://main.zip
	rm main main.zip
