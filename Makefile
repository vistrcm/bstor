.DEFAULT_GOAL := sanity

.PHONY: sanity
sanity: fmt test lint

.PHONY: fmt
fmt:
	go fmt ./... && goimports -w .

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run --enable-all
