THIS_FILE := $(lastword $(MAKEFILE_LIST))
THIS_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: build
build: clean gen
	@sh -c "'$(THIS_DIR)/scripts/build.sh'"

build-api:
	@sh -c "'$(THIS_DIR)/scripts/build.sh' -p thredl.it-api"

clean:
	rm -rf ./bin

build-ui: gen
	@sh -c "'$(THIS_DIR)/scripts/build.sh' -p thredl.it-ui"

build-all: build-api build-ui

gen:
	@sh -c "'$(THIS_DIR)/scripts/gen.sh'"

dev: build
	@sh -c "'$(THIS_DIR)/scripts/dev.sh'"

.PHONY: local
local:
	@echo $(THIS_DIR)
	@sh -c "'$(THIS_DIR)/scripts/local/install.sh'"

.PHONY: local-uninstall
local-uninstall:
	@sh -c "'$(THIS_DIR)/scripts/local/uninstall.sh'"
