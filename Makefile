IMAGE_NAME = platform-slack-invite

GO ?= go
PROTOC ?= protoc

PROTOCFLAGS = --go_out=plugins=grpc:.
PROTOCGATEWAYFLAGS = --grpc-gateway_out=logtostderr=true:.
PROTOSPATH = protos
PLATFORMPROTOPATH = $(GOPATH)/src
GOOGLEAPIPATH = $(GOPATH)/src/github.com/gengo/grpc-gateway/third_party/googleapis
PROTOS = protos/*.proto

DOCKER_REGISTRY = {docker-registry-host:docker-registry-ip}
BUILD_NUMBER = 2

all: proto

proto:
	@echo "Building $(IMAGE_NAME) protocol buffers..."
	$(PROTOC) -I/usr/local/include -I. -I $(PROTOSPATH) -I $(PLATFORMPROTOPATH) -I $(GOOGLEAPIPATH) $(PROTOS) $(PROTOCFLAGS)
	$(PROTOC) -I/usr/local/include -I. -I $(PROTOSPATH) -I $(PLATFORMPROTOPATH) -I $(GOOGLEAPIPATH) $(PROTOS) $(PROTOCGATEWAYFLAGS)

build-amd64:
	env GOOS=linux GOARCH=amd64 go build -v -o build/$(IMAGE_NAME)

docker-build:
	docker build --no-cache -t ${@:2} $(IMAGE_NAME) .

docker-amd64:
	env GOOS=linux GOARCH=amd64 go build -v -o build/$(IMAGE_NAME)
	docker build --no-cache -t ${@:2} $(IMAGE_NAME) .
	docker tag -f $(IMAGE_NAME) $(DOCKER_REGISTRY)/golanghr/$(IMAGE_NAME):$(BUILD_NUMBER)
	docker push $(DOCKER_REGISTRY)/golanghr/$(IMAGE_NAME):$(BUILD_NUMBER)

test:
	$(GO) test

run:
	@echo "Building $(IMAGE_NAME) and running service..."
	$(GO) build && ./$(NAME)
