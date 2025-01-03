test:
    go test ./...

lint:
    golangci-lint run --timeout 5m