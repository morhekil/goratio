include $(dir $(lastword $(MAKEFILE_LIST)))../Makefile.common

CMD_DIR=$(ROOT)
TEST_PACKAGE=github.com/morhekil/goratio/feeder/...

all: get test build

get:
	go get
	go get -t $(TEST_PACKAGE)

test:
	go test $(TEST_PACKAGE)

build:
	go build -o $(DIST_DIR)/$(BINARY)
