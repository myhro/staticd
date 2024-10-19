BINARY ?= staticd
VERSION ?= $(shell ./scripts/version.sh)

build:
	mkdir -p dist/
	CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=$(VERSION)" -o dist/$(BINARY)

clean:
	rm -rf dist/

golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3

gotestsum:
	go install gotest.tools/gotestsum@v1.10.0

lint:
	golangci-lint run

test:
	gotestsum

upx:
	upx dist/*

yamllint:
	yamllint --format colored --strict .github/workflows/ .golangci.yml .yamllint.yml
