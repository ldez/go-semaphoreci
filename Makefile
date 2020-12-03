.PHONY: clean check test

export GO111MODULE=on

default: clean check test

test: clean
	go test -v -cover ./...

clean:
	rm -f cover.out

check:
	golangci-lint run
