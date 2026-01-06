.DEFAULT_GOAL = check

.PHONY: check
check: test lint

.PHONY: test
test:
	@go test ./... -race -count=1

.PHONY: lint
lint:
	@golangci-lint run

