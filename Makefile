.PHONY: all

GOFILES := $(shell go list -f '{{range $$index, $$element := .GoFiles}}{{$$.Dir}}/{{$$element}}{{"\n"}}{{end}}' ./... | grep -v '/vendor/')

default: clean check test

test: clean
	go test -v -cover ./...

dependencies:
	dep ensure -v

clean:
	rm -f cover.out

build:
	go build

fmt:
	@gofmt -s -l -w $(GOFILES)

imports:
	@goimports -w $(GOFILES)

check:
	golangci-lint run
