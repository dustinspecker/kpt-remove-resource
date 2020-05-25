.DEFAULT_GOAL := test

.PHONY: all
all: test-unit build

.PHONY: build
build:
	docker build . --tag kpt-remove-resource:latest

.PHONY: install-kpt
install-kpt:
	./scripts/install-kpt.sh

.PHONY: test
test: test-unit test-integration

.PHONY: test-unit
test-unit:
	go test ./... -cover -coverprofile=cover.out

.PHONY: test-integration
test-integration: build install-kpt
	./tests/integration.sh
