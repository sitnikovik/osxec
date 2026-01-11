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
		go test -race -count=1 $$pkgs; \
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



.PHONY: test-coverage
test-coverage:
	@echo 🧪 Calculating test coverage...
	@mkdir -p tmp
	@pkgs=$$(go list ./... | grep -v 'internal/test/integration$$'); \
	if [ -n "$$pkgs" ]; then \
		go test \
		-covermode=count \
		-coverprofile=tmp/coverage_unit.out \
		$$pkgs \
		>/dev/null; \
		if [ -f tmp/coverage_unit.out ] && [ $$(wc -l < tmp/coverage_unit.out) -le 1 ]; then \
			rm -f tmp/coverage_unit.out; \
		fi; \
	fi
	@go test \
		-tags=integration \
		-coverpkg=./... \
		-covermode=count \
		-coverprofile=tmp/coverage_integration.out \
		./internal/test/integration/... \
		>/dev/null; \
	if [ -f tmp/coverage_integration.out ] && [ $$(wc -l < tmp/coverage_integration.out) -le 1 ]; then \
		rm -f tmp/coverage_integration.out; \
	fi
	@echo "mode: count" > tmp/coverage.out; \
	for f in tmp/coverage_unit.out tmp/coverage_integration.out; do \
		if [ -f "$$f" ]; then \
			tail -n +2 $$f >> tmp/coverage.out; \
		fi; \
	done
	@percent=$$(go tool cover -func=tmp/coverage.out 2>/dev/null | tail -n1 | awk '{print $$NF}'); \
	echo "Total test coverage: $$percent"; \
	percent_no_pct=$${percent%\%}; \
	printf '%s' "$$percent_no_pct" > tmp/coverage_total.out; \
	awk 'NR>1 {n=NF; if($$n==0){ split($$1,a,":"); file=a[1]; split(a[2],b,","); split(b[1],c,"\\."); start=c[1]; split(b[2],d,"\\."); end=d[1]; print file ":" start "-" end } }' tmp/coverage.out \
		| sed 's#^github.com/sitnikovik/osxec/##' \
		| sort -u > tmp/uncovered.out; \