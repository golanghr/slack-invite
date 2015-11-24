NAME = slack-invite

GO ?= go
PROTOC ?= protoc

PROTOCFLAGS = --go_out=plugins=grpc:.
PROTOCGATEWAYFLAGS = --grpc-gateway_out=logtostderr=true:.
PROTOSPATH = protos
PLATFORMPROTOPATH = $(GOPATH)/src
GOOGLEAPIPATH = $(GOPATH)/src/github.com/gengo/grpc-gateway/third_party/googleapis
PROTOS = protos/*.proto

DOCKER_REGISTRY = docker-registry-addr:ip
DOCKER_IMAGE_NAME = platform-slack-invite
BUILD_NUMBER = 2

all: proto

proto:
	@echo "Building $(NAME) protocol buffers..."
	$(PROTOC) -I/usr/local/include -I. -I $(PROTOSPATH) -I $(PLATFORMPROTOPATH) -I $(GOOGLEAPIPATH) $(PROTOS) $(PROTOCFLAGS)
	$(PROTOC) -I/usr/local/include -I. -I $(PROTOSPATH) -I $(PLATFORMPROTOPATH) -I $(GOOGLEAPIPATH) $(PROTOS) $(PROTOCGATEWAYFLAGS)

docker-amd64:
	env GOOS=linux GOARCH=amd64 go build -v -o build/$(DOCKER_IMAGE_NAME)
	docker build --no-cache -t ${@:2} $(DOCKER_IMAGE_NAME) .
	docker tag -f $(DOCKER_IMAGE_NAME) $(DOCKER_REGISTRY)/golanghr/$(DOCKER_IMAGE_NAME):$(BUILD_NUMBER)
	docker push $(DOCKER_REGISTRY)/golanghr/$(DOCKER_IMAGE_NAME):$(BUILD_NUMBER)

test:
	$(GO) test

run:
	@echo "Building $(NAME) and running service..."
	$(GO) build && ./$(NAME)
