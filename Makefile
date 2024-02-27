.PHONY: build test clean

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOFLAGS=-ldflags="-s -w"
BINARY_NAME=sre
BINARY_FOLDER=bin

all: test build

build:
	$(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_NAME) $(GOFLAGS) -v cmd/main.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_FOLDER)/$(BINARY_NAME)
