PHONY: test lint

lint:
	golangci-lint run ./...

test:
	go test -v --race ./...