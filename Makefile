# Run tests for the package.
.PHONY: test
test: go test -v ./tests

# Build the package.
.PHONY: buid
build: go build -v ./...