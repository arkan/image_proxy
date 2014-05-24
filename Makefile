VERSION := 0.1
SHA := $(shell git rev-parse --short HEAD)
CURRENT_DATETIME := $(shell date -u +%d-%m-%Y.%Hh%M)

all: run

build:
	go build -o bin/image_proxy -ldflags "-X main.Version $(VERSION)dev-$(SHA)($(CURRENT_DATETIME))" *.go

build-dist: godep
	godep go build -o bin/image_proxy -ldflags "-X main.Version $(VERSION)dev-$(SHA)($(CURRENT_DATETIME))" *.go

godep:
	go get -u github.com/tools/godep

freeze: deps
	godep save ./...

run:
	go run *.go

deps:
	go get -u -t -v ./...

clean:
	rm -f bin/image_proxy
