.PHONY: test, build, install

test:
	go test ./...

build:
	go build -o splitbbq main.go

install:
	go install
