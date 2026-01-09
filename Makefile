.DEFAULT_GOAL = check

.PHONY: check
check: test lint

.PHONY: test
test:
	@$(MAKE) unit-test || { \
		echo "❌ Unit tests failed; skipping integration tests"; \
		exit 1; \
	}
	@$(MAKE) integration-test

.PHONY: unit-test
unit-test:
	@echo 🧪 Running unit tests...
	@pkgs=$$(go list ./... | grep -v './internal/test/integration$$'); \
	if [ -n "$$pkgs" ]; then \
		go test $$pkgs; \
	else \
		echo "no packages to test"; \
	fi

.PHONY: integration-test
integration-test:
	@echo 🧪 Running integration tests...
	@go test -tags=integration ./internal/test/integration/... -race -count=1

.PHONY: lint
lint:
	@golangci-lint run
