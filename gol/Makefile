# Makefile for the 'gol' Go project in sub-repo 'gol'

# Variables
BINARY := gol
GOBUILD := go build
GOCLEAN := go clean
GOTEST := go test
GOGET := go get
CURDIR := $(shell pwd)
SRCDIR := $(CURDIR)/cmd/

# Targets

all: build
# build compiles the binary
build:
	cd $(SRCDIR) && $(GOBUILD) -o $(CURDIR)/$(BINARY) -v
# test runs the tests
test:
	$(GOTEST) -v ./...
# clean removes the binary
clean:
	$(GOCLEAN)
# run builds and runs the binary
run: build
	$(CURDIR)/$(BINARY) -server=$(CURDIR)/config/server.yaml -db=$(CURDIR)/config/db.yaml
# deps gets the dependencies
deps:
	$(GOGET) ./...
# docs generates the documentation which can be viewed at localhost:8080
# NOTE: package pkgsite should be installed
docs:
	pkgsite
.PHONY: all build test clean run deps docs
