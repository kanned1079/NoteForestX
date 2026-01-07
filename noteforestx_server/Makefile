# Makefile

BINARY_NAME := noteforestx-server

BIN_DIR := ./bin
LDFLAGS := -s -w
MAIN := command/main.go

.PHONY: all build clean \
        build-linux-amd64 \
        build-windows-amd64 \
        build-macos-arm64

all: build

# 默认构建（Linux amd64）
build: build-linux-amd64

build-linux-amd64:
	@echo "Building $(BINARY_NAME) for linux/amd64..."
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 \
	go build -ldflags "$(LDFLAGS)" \
	-o $(BIN_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN)
	@echo "✔ $(BIN_DIR)/$(BINARY_NAME)-linux-amd64"

build-windows-amd64:
	@echo "Building $(BINARY_NAME) for windows/amd64..."
	@mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 \
	go build -ldflags "$(LDFLAGS)" \
	-o $(BIN_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN)
	@echo "✔ $(BIN_DIR)/$(BINARY_NAME)-windows-amd64.exe"

build-darwin-arm64:
	@echo "Building $(BINARY_NAME) for darwin/arm64..."
	@mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=arm64 \
	go build -ldflags "$(LDFLAGS)" \
	-o $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN)
	@echo "✔ $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64"

clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)