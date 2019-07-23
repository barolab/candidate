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

# This version-strategy uses git tags to set the version string
GIT_TAG := $(shell git describe --tags --always --dirty || echo unsupported)
GIT_COMMIT := $(shell git rev-parse --short HEAD || echo unsupported)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
BUILDTIME := $(shell date '+%d-%m-%Y-%Z-%T')

.PHONY: fmt test cover lint install hooks help
default: help

## Build the docker image
build:
	@docker build \
		--build-arg COMMITHASH="${GIT_COMMIT}" \
		--build-arg BUILDTIME="${BUILDTIME}" \
		--build-arg VERSION="${GIT_TAG}" \
		--tag "${ORG}/${BIN}" \
		.

## Format go source code
fmt:
	@go fmt ./...

## Execute unit tests
test:
	@go test ./...

## Compute coverage
cover:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

## Install dependencies used for development
install: hooks
	@go mod vendor

## Install git hooks for post-checkout & pre-commit
hooks:
	@cp -f ./scripts/post-checkout .git/hooks
	@cp -f ./scripts/pre-commit .git/hooks
	@chmod +x .git/hooks/post-checkout
	@chmod +x .git/hooks/pre-commit

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
