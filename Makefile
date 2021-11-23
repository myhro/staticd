BINARY ?= staticd
VERSION ?= $(shell ./scripts/version.sh)

build:
	mkdir -p dist/
	go build -ldflags="-s -w -X main.version=$(VERSION)" -o dist/$(BINARY)

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

yamllint:
	yamllint --format colored --strict .github/workflows/ .golangci.yml .yamllint.yml
