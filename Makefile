PROJECTNAME := spacebotcalc
GOPATH := $(shell go env GOPATH)
VERSION := $(shell cat VERSION)
BUILD := $(shell git rev-parse --short HEAD)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION) -X=main.build=$(BUILD)"

.PHONY: build install run clean help

all: build

## build: Compile the binary.
build:
	go build $(LDFLAGS) -o $(PROJECTNAME) cmd/$(PROJECTNAME)/main.go

## install: Install to $GOBIN path.
install: build
	install $(PROJECTNAME) $(GOPATH)/bin

## run: Run code.
run:
	go run cmd/$(PROJECTNAME)/main.go

## clean: Cleanup binary.
clean:
	-@rm -f $(PROJECTNAME)

## help: Show this message.
help: Makefile
	@echo "Available targets:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
