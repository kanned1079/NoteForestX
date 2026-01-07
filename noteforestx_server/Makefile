# Makefile

BINARY_NAME := noteforestx-server

GOOS := linux
GOARCH := amd64

BIN_DIR := ./bin

LDFLAGS := -s -w

.PHONY: all build clean

all: build

build:
	@echo "Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(BIN_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY_NAME) command/main.go
	@echo "Build finished: $(BIN_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)