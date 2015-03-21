.PHONY: build doc fmt lint run test vendor_clean vendor_get vendor_update vet

# Prepend our _vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
GOPATH := ${PWD}/_vendor:${GOPATH}
BREW_PREFIX := $(shell brew --prefix 2>/dev/null)
GO := $(BREW_PREFIX)/bin/go
GOROOT := $(BREW_PREFIX)/Cellar/go/1.4.2/libexec
CGO_CFLAGS := -I$(BREW_PREFIX)/include
CGO_LDFLAGS := -L$(BREW_PREFIX)/lib
export CGO_CFLAGS CGO_LDFLAGS GOROOT GOPATH

default: vet run

bootstrap: vendor_update

run: build
	bin/gtm list

build:
	/usr/bin/env CC=clang \
		$(GO) build -v -o ./bin/gtm \
		./src/gtm.go ./src/client.go

fmt:
	$(GO) fmt ./src/...

test: vet
	$(GO) test ./src/...

vendor_clean:
	rm -dRf ./_vendor/src

vendor_get: vendor_clean
	GOPATH=${PWD}/_vendor $(GO) get -d -u -v \
	github.com/codegangsta/cli \
	github.com/octokit/go-octokit/octokit \
	golang.org/x/tools/cmd/vet

vendor_update: vendor_get
	rm -rf `find ./_vendor/src -type d -name .git` \
	&& rm -rf `find ./_vendor/src -type d -name .hg` \
	&& rm -rf `find ./_vendor/src -type d -name .bzr` \
	&& rm -rf `find ./_vendor/src -type d -name .svn`

vet:
	$(GO) vet ./src/...
