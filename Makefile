.PHONY: build clean

VERSION=$(shell git rev-parse HEAD)

build: cgame-linux-amd64 cgame-linux-arm cgame-darwin-amd64 cgame-windows-amd64.exe cgame-windows-386.exe

.get-deps: *.go
	go get -t -d -v ./...
	touch .get-deps

clean:
	rm -f .get-deps
	rm -f *-amd64* *-386*

test:
	go test

cgame-linux-amd64: .get-deps *.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ *.go

cgame-linux-arm: .get-deps *.go
	GOOS=linux GOARCH=arm go build -ldflags "-X main.version=$(VERSION)" -o $@ *.go

cgame-darwin-amd64: .get-deps *.go
	GOOS=darwin go build -ldflags "-X main.version=$(VERSION)" -o $@ *.go

cgame-windows-amd64.exe: .get-deps *.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ *.go

cgame-windows-386.exe: .get-deps *.go
	GOOS=windows GOARCH=386 go build -ldflags "-X main.version=$(VERSION)" -o $@ *.go