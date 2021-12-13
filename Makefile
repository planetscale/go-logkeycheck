lint:
	golangci-lint run -v ./...

test-deps:
	cd internal/logkeycheck/testdata && make src

test: test-deps
	go test -cover -v ./...

build:
	go build .

build-plugin:
	CGO_ENABLED=1 go build -buildmode=plugin .

.PHONY: lint test build build-plugin