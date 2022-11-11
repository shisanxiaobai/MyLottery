
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/micro/micro/v2/cmd/protoc-gen-micro@latest
	go install github.com/micro/micro/v2/cmd/protoc-gen-openapi@latest


	
.PHONY: build
build:
	go build -o lottery-gateway *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t lottery-gateway:latest
