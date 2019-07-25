# COLORS
TARGET_MAX_CHAR_NUM := 10
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

# The binary to build (just the basename).
PWD := $(shell pwd)
NOW := $(shell date +%s)
BIN := candidate
ORG := barolab
PKG := github.com/${ORG}/${BIN}

GO ?= go
GOFMT ?= gofmt -s
GOFILES := $(shell find . -name "*.go" -type f)
PACKAGES ?= $(shell $(GO) list ./...)
NAME ?= candidate

# This version-strategy uses git tags to set the version string
GIT_TAG := $(shell git describe --tags --always --dirty || echo unsupported)
GIT_COMMIT := $(shell git rev-parse --short HEAD || echo unsupported)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
BUILDTIME := $(shell date '+%d-%m-%Y-%Z-%T')

.PHONY: fmt fmt-check vet test test-coverage cover install hooks example help
default: help

## Format go source code
fmt:
	$(GOFMT) -w $(GOFILES)

## Check if source code is formatted correctly
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

## Check source code for common errors
vet:
	$(GO) vet ${PACKAGES}

## Execute unit tests
test:
	$(GO) test ${PACKAGES}

## Execute unit tests & compute coverage
test-coverage:
	$(GO) test -coverprofile=coverage.out ${PACKAGES}

## Compute coverage
cover: test-coverage
	$(GO) tool cover -html=coverage.out

## Install dependencies used for development
install: hooks
	$(GO) mod download

## Install git hooks for post-checkout & pre-commit
hooks:
	@cp -f ./scripts/post-checkout .git/hooks
	@cp -f ./scripts/pre-commit .git/hooks
	@chmod +x .git/hooks/post-checkout
	@chmod +x .git/hooks/pre-commit

## Run the example
example:
	@go run example/main.go ${NAME}

## Print this help message
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
