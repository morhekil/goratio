include $(dir $(lastword $(MAKEFILE_LIST)))../Makefile.common

TEST_PACKAGE=github.com/morhekil/goratio/analyser/...
BINARY=goratio-analyser

all: get test build

get:
	go get
	go get -t $(TEST_PACKAGE)

test:
	go test $(TEST_PACKAGE)

build:
	go build -o $(DIST_DIR)/$(BINARY)
