.PHONY: dev run build tidy clean test

## dev: start server with live-reload (air)
dev:
	air

## run: start server without live-reload
run:
	go run ./cmd/server

## build: compile binary to ./bin/server
build:
	go build -o ./bin/server.exe ./cmd/server

## tidy: sync go.mod / go.sum
tidy:
	go mod tidy

## test: run all tests with verbose output
test:
	go test ./... -v

## clean: remove build artifacts
clean:
	rm -rf ./tmp ./bin
