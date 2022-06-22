TARGETS           ?= linux/amd64 darwin/arm64
PROJECT_NAME	  := router-sidecar
PKG				  := airnity.com/$(PROJECT_NAME)

# go option
GO        ?= go
# Uncomment to enable vendor
GO_VENDOR := # -mod=vendor
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   := -w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
DISTDIR   := dist

# Required for globs to work correctly
SHELL=/usr/bin/env bash

#  Version

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
DATE	   = $(shell date +%F_%T%Z)

BINARY_VERSION = ${GIT_SHA}

HAS_GORELEASER := $(shell command -v goreleaser;)
HAS_GIT := $(shell command -v git;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_CURL:=$(shell command -v curl;)
HAS_MOCKGEN:=$(shell command -v mockgen;)
# Uncomment to use gox instead of goreleaser
HAS_GOX := $(shell command -v gox;)

.DEFAULT_GOAL := code/lint

#############
#   Build   #
#############

.PHONY: code/build
code/build: code/clean setup/dep/install
	$(GO) build $(GO_VENDOR) -o $(BINDIR)/$(PROJECT_NAME) $(GOFLAGS) -tags '$(TAGS)' $(PKG)/cmd/${PROJECT_NAME}

# Uncomment to use gox instead of goreleaser
.PHONY: code/build-cross
code/build-cross: code/clean setup/dep/install
	CGO_ENABLED=0 GOFLAGS="-trimpath $(GO_VENDOR)" gox -output="$(DISTDIR)/bin/{{.OS}}-{{.Arch}}/{{.Dir}}" -osarch='$(TARGETS)' $(if $(TAGS),-tags '$(TAGS)',) ${PKG}/cmd/${PROJECT_NAME}

# .PHONY: code/build-cross
# code/build-cross: code/clean setup/dep/install
# ifdef HAS_GORELEASER
# 	goreleaser --snapshot --skip-publish
# endif
# ifndef HAS_GORELEASER
# 	curl -sL https://git.io/goreleaser | bash -s -- --snapshot --skip-publish
# endif

.PHONY: code/clean
code/clean:
	@rm -rf $(BINDIR) $(DISTDIR)

#############
#  Release  #
#############

# .PHONY: release/all
# release/all: code/clean setup/dep/install
# ifdef HAS_GORELEASER
# 	goreleaser
# endif
# ifndef HAS_GORELEASER
# 	curl -sL https://git.io/goreleaser | bash
# endif

#############
#   Tests   #
#############

.PHONY: test/all
test/all: setup/dep/install
	$(GO) test $(GO_VENDOR) --tags=unit,integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/unit
test/unit: setup/dep/install
	$(GO) test $(GO_VENDOR) --tags=unit -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/integration
test/integration: setup/dep/install
	$(GO) test $(GO_VENDOR) --tags=integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/coverage
test/coverage:
	cat c.out.tmp | grep -v "mock_" > c.out
	$(GO) tool cover -html=c.out -o coverage.html
	$(GO) tool cover -func c.out

#############
#   Setup   #
#############
.PHONY: setup/dep/install
setup/dep/install:
ifndef HAS_GOLANGCI_LINT
	@echo "=> Installing golangci-lint tool"
ifndef HAS_CURL
	$(error You must install curl)
endif
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.42.1
endif
ifndef HAS_GIT
	$(error You must install Git)
endif
ifndef HAS_MOCKGEN
	@echo "=> Installing mockgen tool"
	$(GO) get -u github.com/golang/mock/mockgen@v1.6.0
	$(GO) mod tidy
endif
# Uncomment to use gox instead of goreleaser
ifndef HAS_GOX
	@echo "=> Installing gox"
	$(GO) install github.com/mitchellh/gox@latest
	$(GO) mod tidy
endif
	$(GO) mod download

.PHONY: setup/dep/tidy
setup/dep/tidy:
	$(GO) mod tidy

.PHONY: setup/dep/update
setup/dep/update:
	$(GO) get -u ./...

.PHONY: setup/dep/vendor
setup/dep/vendor:
	$(GO) mod vendor
