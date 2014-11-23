include Makefile.common

all: $(DIST_DIR)/$(BINARY)

$(DIST_DIR)/$(BINARY):
	cowsay "Building feeder binary"
	@$(MAKE) -C feeder/cmd all

clean:
	rm -f dist/*
