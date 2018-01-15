.DEFAULT_GOAL := build

COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE = `date +%FT%T%z`

GO = go
BINARY_DIR=bin

BUILD_DEPS:= github.com/alecthomas/gometalinter

.PHONY: test build

help:
	@echo "build      - go build"
	@echo "test       - go test"
	@echo "checkstyle - gofmt+golint+misspell"


get-build-deps:
	$(GO) get $(BUILD_DEPS)
	gometalinter --install

test:
	$(GO) test -v

checkstyle:
	gometalinter --vendor ./... --fast --disable=gas --disable=errcheck --disable=gotype --deadline 10m

fmt:
	gofmt -l -w -s .

# Builds the project
build: checkstyle test
	$(GO) build .


clean:
	if [ -d ${BINARY_DIR} ] ; then rm -r ${BINARY_DIR} ; fi

release:
	git tag -a ${v} -m "creating tag ${v}"
	git push origin "refs/tags/${v}"
