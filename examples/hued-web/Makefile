GOROOT ?= /usr/local/go
GOBIN ?= go
GOPATH = $(shell pwd)/lib
GODEPS = github.com/golang/glog
HUED_WEB_BIN = bin
HUED_WEB_SOURCE = hued-web.go

HUED_WEB_APPS = $(HUED_WEB_BIN)/hued-web-arm $(HUED_WEB_BIN)/hued-web

all: $(HUED_WEB_APPS)

$(HUED_WEB_BIN)/hued-web-arm: deps
	GOPATH=$(GOPATH) GOARCH=arm GOARM=5 GOOS=linux $(GOBIN) build $(HUED_WEB_SOURCE)
	mv hued-web hued-web-arm
	mv hued-web-arm $(HUED_WEB_BIN)

$(HUED_WEB_BIN)/hued-web: deps
	GOPATH=$(GOPATH) $(GOBIN) build $(HUED_WEB_SOURCE)
	mv hued-web $(HUED_WEB_BIN)

$(GOPATH)/%:
	GOPATH=$(GOPATH) $(GOBIN) get $*

deps: fix-gopath $(patsubst %, $(GOPATH)/%, $(GODEPS))

fix-gopath:
	[ -d lib ] || mkdir lib
	[ -d bin ] || mkdir bin

clean:
	rm -rf lib bin