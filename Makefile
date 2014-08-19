GOROOT ?= /usr/local/go
GOBIN ?= go
GOPATH = $(shell pwd)/lib
GODEPS = 
HUE_SRC = src
HUE_VERSION ?= 1.2.0
HUE_INSTALL ?= /usr/local/go-hue
HUE_INSTALL_FULL ?= $(HUE_INSTALL)/$(HUE_VERSION)/$(HUE_BIN)
HUE_SOURCES_TEST = $(HUE_SRC)/configuration \
$(HUE_SRC)/groups \
$(HUE_SRC)/portal \
$(HUE_SRC)/lights

all: $(HUE_APPS)

$(GOPATH)/src/%:
	GOPATH=$(GOPATH) $(GOBIN) get $*

%_test.go:
	GOPATH=$(GOPATH) go test $<

test: deps
	$(foreach var,$(HUE_SOURCES_TEST), pushd $(var); GOPATH=$(GOPATH) $(GOBIN) test || exit 1; popd;)

deps: fix-gopath $(patsubst %, $(GOPATH)/src/%, $(GODEPS))

fix-gopath:
	[ -d lib ] || mkdir lib

clean:
	rm -rf lib $(HUE_INSTALL)/$(HUE_VERSION)
