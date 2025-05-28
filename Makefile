.PHONY: test cover

# Run all tests
test:
	go test ./...

# Run tests with coverage
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Clean coverage artifacts
clean:
	rm -f coverage.out
