default_target: usage
.PHONY : default_target upload

# VERSION := $(shell $(PYTHON) setup.py --version)

usage:
	@echo "The blowpipe makefile"
	@echo ""
	@echo "Usage : make <command> "
	@echo ""
	@echo " commands"
	@echo "  setup                 - install tooling dependencies"
	@echo "  proto                 - rebuilds protobuf/grpc files"
	@echo "  test                  - runs tests"
	@echo "  build                 - builds executable (no tests)"
	@echo "  all                   - runs tests and build executable"
	@echo ""

setup:
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger 
	go get github.com/golang/protobuf/protoc-gen-go
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger 
	go install github.com/golang/protobuf/protoc-gen-go

clean:
	go clean

test:
	go test

proto:
	protoc -I ../../grpc blowpiperpc.proto --go_out=plugins=grpc:blowpiperpc

build:
	go build

all: proto clean test build
	go fmt

install: build
	cd blowpipe && go install
