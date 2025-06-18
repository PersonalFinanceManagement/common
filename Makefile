.PHONY: all fmt test tidy vet clean

all: fmt vet test

fmt:
	@echo "Running go fmt..."
	go fmt ./...

tidy:
	@echo "Running go mod tidy..."
	go mod tidy
	go mod vendor

test:
	@echo "Running go test..."
	go test -v ./...

vet:
	@echo "Running go vet..."
	go vet ./...

clean:
	@echo "Cleaning up..."
	go clean -modcache
	rm -rf ./vendor
	rm -f go.sum
	rm -f go.mod
	rm -f go.work

build:
	@echo "Building the project..."
	go build -o bin/ ./...

help:
	@echo "Available commands:"
	@echo "  make all    - Run fmt, vet, and test"
	@echo "  make fmt    - Format the code"
	@echo "  make tidy   - Tidy up go modules"
	@echo "  make test   - Run tests"
	@echo "  make vet    - Run go vet"
	@echo "  make clean  - Clean up build artifacts and modules"
	@echo "  make build  - Build the project"
	@echo "  make help   - Show this help message"