build:
	mkdir -p dist/
	go build -o dist/staticd

clean:
	rm -rf dist/

test:
	go test -v ./...
