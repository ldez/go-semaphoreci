.PHONY: clean check test build fmt imports

export GO111MODULE=on

GOFILES := $(shell go list -f '{{range $$index, $$element := .GoFiles}}{{$$.Dir}}/{{$$element}}{{"\n"}}{{end}}' ./... | grep -v '/vendor/')

default: clean check test

test: clean
	go test -v -cover ./...

clean:
	rm -f cover.out

fmt:
	@gofmt -s -l -w $(GOFILES)

imports:
	@goimports -w $(GOFILES)

check:
	golangci-lint run
