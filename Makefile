.DEFAULT_GOAL := sanity

.PHONY: sanity
sanity: fmt test lint-all

.PHONY: fmt
fmt:
	go fmt ./... && goimports -w .

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-all
lint-all:
	golangci-lint run --enable-all

.PHONY: pre-commit
pre-commit: test lint
