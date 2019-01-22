GOCMD=go
GOARCH=wasm
GOOS=js
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=lib.wasm

all: build
build:
		GOARCH=$(GOARCH) GOOS=$(GOOS) $(GOBUILD) -o $(BINARY_NAME) -v