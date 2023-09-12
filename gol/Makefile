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

build:
	cd $(SRCDIR) && $(GOBUILD) -o $(CURDIR)/$(BINARY) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)

run: build
	$(CURDIR)/$(BINARY) -server=$(CURDIR)/config/server.yaml -db=$(CURDIR)/config/db.yaml

deps:
	$(GOGET) ./...

docs:
	pkgsite
.PHONY: all build test clean run deps docs
