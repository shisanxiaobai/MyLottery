
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/micro/micro/v2/cmd/protoc-gen-micro@latest
	go install github.com/micro/micro/v2/cmd/protoc-gen-openapi@latest

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/lottery-user.proto
	protoc --openapi_out=. --proto_path=. proto/lottery-admin.proto
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/user/user.proto
	protoc --proto_path=. --micro_out=. --go_out=:. proto/admin/admin.proto

.PHONY: build
build:
	go build -o lottery-user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t lottery-user:latest
