.DEFAULT_GOAL := build

COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE = `date +%FT%T%z`

GO = go
BINARY_DIR=bin
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PACKAGES_NOVENDOR = $(shell glide novendor)

BUILD_DEPS:= github.com/alecthomas/gometalinter

.PHONY: vendor test build

help:
	@echo "build      - go build"
	@echo "test       - go test"
	@echo "checkstyle - gofmt+golint+misspell"

vendor:
	## Requires glide!
	glide install

get-build-deps:
	$(GO) get $(BUILD_DEPS)
	gometalinter --install

test:
	$(GO) test -v

checkstyle:
	gometalinter --vendor ./... --fast --disable=gas --disable=errcheck --disable=gotype --deadline 10m

fmt:
	gofmt -l -w -s ${GOFILES_NOVENDOR}

# Builds the project
build: checkstyle test
	$(GO) build $(PACKAGES_NOVENDOR)


clean:
	if [ -d ${BINARY_DIR} ] ; then rm -r ${BINARY_DIR} ; fi

release:
	git tag -a ${v} -m "creating tag ${v}"
	git push origin "refs/tags/${v}"
