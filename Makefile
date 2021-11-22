BINARY ?= staticd

build:
	mkdir -p dist/
	go build -ldflags="-s -w" -o dist/$(BINARY)

clean:
	rm -rf dist/

test:
	go test -v ./...

upx:
	upx dist/*
