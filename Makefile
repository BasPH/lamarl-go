GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GIT_COMMIT=$(shell git rev-parse --short HEAD)

.PHONY: all fmt vet lint test build update-lambda dockerize

all: fmt vet lint test build update-lambda

fmt:
	gofmt -s -l -w $(GOFILES)

vet:
	go tool vet $(GOFILES)

lint:
	for f in $(GOFILES); do golint $f; done

build:
	GOOS=linux go build -o main cmd/lambda/*
	go build -o sushigo-server cmd/sushigo/*

test:
	go test ./...

update-lambda:
	zip main.zip ./main
	aws lambda update-function-code \
    	--function-name awscodestar-lamarl-go-lambda-GetHelloWorld-1KA3SG4UKTUCO \
    	--zip-file fileb://main.zip
	rm main main.zip

dockerize:
	@echo ">> dockerizing"
	docker rm -f $(docker ps -aq) || true
	go-bindata-assetfs -o static_bindata.go -pkg static static/...
	go-bindata -o templates/template_bindata.go -pkg templates templates/...
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sushigo-server cmd/sushigo/*
	docker build -t sushigo-server -f Dockerfile .
	docker run -d -p 8080:8080 sushigo-server
