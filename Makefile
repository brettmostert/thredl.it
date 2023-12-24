THIS_FILE := $(lastword $(MAKEFILE_LIST))
THIS_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: build
build:
	@sh -c "'$(THIS_DIR)/scripts/build.sh'"

.PHONY: local
local:
	@echo $(THIS_DIR)
	@sh -c "'$(THIS_DIR)/scripts/local/install.sh'"

.PHONY: local-uninstall
local-uninstall:
	@sh -c "'$(THIS_DIR)/scripts/local/uninstall.sh'"
