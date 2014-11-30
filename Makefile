include Makefile.common

all: $(DIST_DIR)/goratio-feeder $(DIST_DIR)/goratio-analyser

$(DIST_DIR)/goratio-feeder:
	cowsay "Building feeder binary"
	@$(MAKE) -C feeder all

$(DIST_DIR)/goratio-analyser:
	cowsay "Building analyser binary"
	@$(MAKE) -C analyser all

clean:
	rm -f dist/*
