.PHONY: all
all: test-unit build

.PHONY: build
build:
	docker build . --tag kpt-remove-resource:latest

.PHONY: test-unit
test-unit:
	go test ./... -cover -coverprofile=cover.out
