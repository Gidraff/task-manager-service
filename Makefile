OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

BIN=api
REGISTRY ?= gidraff
VERSION := $(shell git describe --tags --always --dirty)

BUILD_DIR := bin
IMAGE_NAME := taskman

TAG := $(VERSION)__$(OS)_$(ARCH)
IMAGE := $(REGISTRY)/$(IMAGE_NAME):$(TAG)

REPORTS := reports

test: $(REPORTS)
	go test -coverprofile=./$(REPORTS)/coverage.out ./...
	@echo "Running test with cover"

cover: $(REPORTS)
	go tool cover -func=./$(REPORTS)/coverage.out

unittest:
	go test -short ./...

engine: $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BIN) cmd/api/*.go
	@echo "Building $(BINARY)"

image: build
	docker build -t $(IMAGE) .

push: image
	docker push $(IMAGE)

$(BUILD_DIR):
		@mkdir -p $@

$(REPORTS):
	@mkdir -p $@

clean:
	if [ -f ${BIN} ] ; then rm ${BIN} ; fi

bin-clean:
	if [ -d ./bin ]; then rm -rf ./bin ; fi
	echo "Cleaning bin folder"

.PHONY: build
