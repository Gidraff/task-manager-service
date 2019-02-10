# include .env
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMDX) get
BINARY_NAME=serviceapp
BINARY_UNIX=$(BINARY_NAME)_unix

all: clean format build run
	@echo "Starting app"


build:
	@echo "Building application..."
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	@echo "Hello from target one..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	@echo "Done cleaning"

format:
	@echo "Format code"
	$(GOCMD) fmt

run: build
	@echo "Starting application server..."
	./$(BINARY_NAME)
	@echo "Done!"
