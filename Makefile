VERBOSE_ORIGINS := "command line" "environment"
ifdef V
  ifeq ($(filter $(VERBOSE_ORIGINS),$(origin V)),)
    BUILD_VERBOSE := $(V)
  endif
endif

ifndef BUILD_VERBOSE
  BUILD_VERBOSE := 0
endif

ifeq ($(BUILD_VERBOSE),1)
  Q :=
else
  Q := @
endif

PHONY += all
CURDIR := $(shell pwd)
BUILD_DIR ?= $(CURDIR)/build
GOBIN_DIR := $(BUILD_DIR)/bin
GOBIN_EXAMPLE_DIR := $(BUILD_DIR)/bin/examples
HOST_DIR := $(BUILD_DIR)/host
HOSTBIN_DIR := $(HOST_DIR)/bin
GOTOOLSBIN_DIR := $(HOSTBIN_DIR)
GOTOOLS_DIR := $(CURDIR)/tools
TMP_DIR := $(BUILD_DIR)/tmp
DIRS := \
	$(GOBIN_DIR) \
	$(HOST_DIR) \
	$(HOSTBIN_DIR) \
	$(TMP_DIR) \
	$(GOBIN_EXAMPLE_DIR)

HOST_OS := $(shell uname -s)

REV ?= $(shell git rev-parse --short HEAD 2> /dev/null)
GOPATH ?= $(shell go env GOPATH)

export PATH:=$(GOTOOLS_DIR):$(HOSTBIN_DIR):$(PATH)
export REV

define find-subdir
$(shell find $(1) -maxdepth 1 -mindepth 1 -type d -o -type l)
endef

APPS := $(sort $(notdir $(call find-subdir,cmd)))
PHONY += $(APPS)

all: $(APPS) examples

$(DIRS) :
	$(Q)mkdir -p $@

EXAMPLES := $(sort $(notdir $(call find-subdir,examples)))

.SECONDEXPANSION:
$(EXAMPLES): $(addprefix $(GOBIN_EXAMPLE_DIR)/,$$@)

.PHONY: examples
examples: $(EXAMPLES)
$(GOBIN_EXAMPLE_DIR)/%: $(GOBIN_DIR) FORCE
	$(Q)go build -o $@ ./examples/$(notdir $@)
	@echo "Done building."
	@echo "Run \"$(subst $(CURDIR),.,$@)\" to launch $(notdir $@)."

.SECONDEXPANSION:
$(APPS): $(addprefix $(GOBIN_DIR)/,$$@)
$(GOBIN_DIR)/%: $(GOBIN_DIR) FORCE
	$(Q)go build -o $@ ./cmd/$(notdir $@)
	@echo "Done building."
	@echo "Run \"$(subst $(CURDIR),.,$@)\" to launch $(notdir $@)."

CODEGEN_DEPS := \
	# $(GOTOOLS_DIR)/abigen

.PHONY: gen
gen: $(CODEGEN_DEPS)
	$(Q)go generate -x ./...

.PHONY: test
test:
	$(Q)go test -v ./...

.PHONY: clean
clean:
	$(Q)rm -fr $(GOBIN_DIR) $(HOST_DIR)

.PHONY: help
help:
	@echo  'Generic targets:'
	@echo  '  all                         - Build all targets marked with [*]'
	@for app in $(APPS); do \
		printf "* %s\n" $$app; done
	@echo  '* examples'
	@echo  ''
	@echo  'Code generation targets:'
	@echo  '  gen                         - Generate Go code from various sources'
	@echo  ''
	@echo  'Examples targets:'
	@echo  '  examples                    - Build all examples'
	@echo  ''
	@echo  'Test targets:'
	@echo  '  test                        - Run all tests'
	@echo  ''
	@echo  'Cleaning targets:'
	@echo  '  clean                       - Remove built executables'
	@echo  ''
	@echo  'Execute "make" or "make all" to build all targets marked with [*] '
	@echo  'For further info see the ./README.md file'

.PHONY: $(PHONY)

.PHONY: FORCE
FORCE: