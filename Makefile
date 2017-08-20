.PHONY: build clean

VERSION=$(shell git rev-parse HEAD)

build: cgame_amd64 cgame_darwin

.get-deps: *.go
	go get -t -d -v ./...
	touch .get-deps

clean:
	rm -f .get-deps
	rm -f *_amd64 *_darwin

cgame_amd64: .get-deps *.go
	 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ *.go

cgame_darwin: .get-deps *.go
	GOOS=darwin go build -ldflags "-X main.version=$(VERSION)" -o $@ *.go
