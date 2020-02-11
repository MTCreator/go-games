# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGENERATE=$(GOCMD) generate

# Directories
COMMANDS_DIR=cmd
BUILD_DIR=bin

# All command files
COMMAND_FILES := $(shell find $(COMMANDS_DIR) -name '*.go')

.PHONY: all clean

all: test build

build: directories generate
	$(GOBUILD) -o $(BUILD_DIR) -v ./cmd/...

build-debug: directories generate
	$(GOBUILD) -race -v ./cmd/...

debug: build-debug

test:
	$(GOTEST) -v ./...

generate:
	$(GOGENERATE) ./...

clean: 
	$(GOCLEAN)

directories:
	mkdir -p bin