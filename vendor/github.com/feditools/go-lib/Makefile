PROJECT_NAME=login

.DEFAULT_GOAL := test

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

lint:
	@echo linting
	@golint $(shell go list ./... | grep -v /vendor/)

test: tidy fmt lint
	go test -cover ./...

tidy:
	go mod tidy

vendor: tidy
	go mod vendor

.PHONY: fmt lint test tidy vendor