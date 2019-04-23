PACKAGES = $(shell go list ./... | grep -v '/vendor/')

PACKAGE = github.com/feeltheajf/go-raml
COMMIT_HASH = $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE = $(shell date +%FT%T%z)
TMP ?= $(GOPATH)/src/tmp

ldflagsversion = -X $(PACKAGE)/commands.CommitHash=$(COMMIT_HASH) -X $(PACKAGE)/commands.BuildDate=$(BUILD_DATE) -s -w

all: install

install: generate
	go install -v -ldflags '$(ldflagsversion)'

build: generate
	go build

generate:
	go generate

test: generate
	$(foreach PKG,$(PACKAGES),go test -v -coverprofile=coverage.xml -covermode=atomic $(PKG);)

gogentest:
	cd codegen/golang/gentest; bash test.sh

pygentest:
	cd codegen/python/gentest; bash test.sh

pyclient: build
	./go-raml client \
		--language python \
		--dir $(TMP)/pyclient \
		--ramlfile $(RAML)

goclient: build
	./go-raml client \
		--language go \
		--dir $(TMP)/goclient \
		--package client \
		--ramlfile $(RAML)
	cd $(TMP)/goclient && go install
	cd $(TMP)/goclient/types && go install

clean:
	rm -rf $(TMP)
