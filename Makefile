BINARY ?= staticd

build:
	mkdir -p dist/
	go build -ldflags="-s -w" -o dist/$(BINARY)

clean:
	rm -rf dist/

golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0

lint:
	golangci-lint run

test:
	go test -v ./...

upx:
	upx dist/*
