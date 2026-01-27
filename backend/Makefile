# SteamInputDB Makefile

BINARY_NAME := steaminputdb
MAIN_PKG := ./cmd/steaminputdb
SRC_DIR := .
DIST_DIR := ./dist

VERSION ?= $(shell git describe --tags --match "v[0-9]*.[0-9]*.[0-9]*" --always 2>/dev/null || echo v0.0.0-dev)
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
BUILD_TIME := $(shell date -u +"%Y-%m-%d_%H:%M:%S")

LDFLAGS := -s -w -X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.Date=$(BUILD_TIME)
BUILD_FLAGS := -trimpath -ldflags "$(LDFLAGS)"

.PHONY: all
all: test build

.PHONY: help
help: ## Show this help message
	@echo SteamInputDB Makefile
	@echo
	@echo Usage: make [target]
	@echo
	@echo Build Targets:
	@echo   build                Build steaminputdb for current platform
	@echo   clean                Remove build artifacts
	@echo   test                 Run tests
	@echo   deps                 Download Go dependencies
	@echo   tidy                 Tidy Go dependencies
	@echo   fmt                  Format Go code
	@echo   vet                  Run go vet
	@echo   lint                 Run golangci-lint
	@echo   run                  Build and run steaminputdb server
	@echo   dev                  Run server with hot reload (air)
	@echo   version              Show version information

.PHONY: deps
deps: 
	cd $(SRC_DIR) && go mod download

.PHONY: tidy
tidy: 
	cd $(SRC_DIR) && go mod tidy

.PHONY: vet
vet: 
	cd $(SRC_DIR) && go vet ./...

.PHONY: test
test: 
	cd $(SRC_DIR) && go test -count=1 -v ./...

.PHONY: clean
clean: 
	rm -rf $(DIST_DIR)
	rm -f $(SRC_DIR)/coverage.out $(SRC_DIR)/coverage.html

.PHONY: fmt
fmt: 
	cd $(SRC_DIR) && go fmt ./...

.PHONY: lint
lint: 
	cd $(SRC_DIR) && golangci-lint run

.PHONY: build
build:
	mkdir -p $(DIST_DIR)
	cd $(SRC_DIR) && go build $(BUILD_FLAGS) -o $(DIST_DIR)/$(BINARY_NAME) $(MAIN_PKG)

.PHONY: run
run: 
	cd $(SRC_DIR) && go run $(MAIN_PKG) server

.PHONY: dev
dev: ## Run server with hot reload using air
	cd $(SRC_DIR) && air

.PHONY: version
version: 
	@echo Version: $(VERSION)
	@echo Commit:  $(COMMIT)
	@echo Built:   $(BUILD_TIME)
