#! /usr/bin/make

all: clean generate

clean:
	@rm -rf ./gen

generate:
	goa gen github.com/sgerogia/hello-goa/design

test:
	go test ./...

keys:
	rm -f /tmp/test_rsa*
	ssh-keygen -q -t rsa -m PEM -f /tmp/test_rsa -q -N ""
	openssl rsa -in /tmp/test_rsa -pubout -out /tmp/test_rsa.pub

build:
	go build ./cmd/server

build-cli:
	go build ./cmd/server-cli

run-local-http: keys build
	./server -debug -private-key /tmp/test_rsa -public-key /tmp/test_rsa.pub



