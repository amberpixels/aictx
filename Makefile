# Variables
GOLANGCI_LINT := $(shell which golangci-lint)

BUILD_DIR := build
CMD_DIR = ./cmd/aictx
MAIN_FILE := $(CMD_DIR)/main.go

BINARY_NAME := aictx
INSTALL_DIR := $(shell go env GOPATH)/bin

# Default target
all: build

# Build the binary
build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# Run the binary
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Tidy: format and vet the code
tidy:
	@go fmt $$(go list ./...)
	@go vet $$(go list ./...)
	@go mod tidy

# Install golangci-lint only if it's not already installed
lint-install:
	@if ! [ -x "$(GOLANGCI_LINT)" ]; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi

# Lint the code using golangci-lint
# todo reuse var if possible
lint: lint-install
	$(shell which golangci-lint) run

# Install the binary globally with aliases
install:
	@go install $(CMD_DIR)

# Uninstall the binary and remove the alias
uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

readme: build tools/readme.template.md
	@go run tools/readmegen.go

# Phony targets
.PHONY: all build run tidy lint-install lint install uninstall readme
