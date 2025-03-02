# XReader Makefile
.PHONY: all build test run clean vet fmt

# Default target
all: build

# Build the XReader executable
build:
	go build -o XReader -v

# Run unit tests
test:
	go test ./internal/... -v

# Run the application
run:
	./XReader --help

# Clean up build artifacts
clean:
	rm -f XReader

# Format code
fmt:
	go fmt ./...

# Static analysis
vet:
	go vet ./...

# Install dependencies
deps:
	go mod tidy