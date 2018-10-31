default: tests

tests: golangci-lint critic unittest

unittest:
	@sh -c "'$(CURDIR)/scripts/gotest.sh'"

golangci-lint:
	@sh -c "'$(CURDIR)/scripts/golangci_lint_check.sh'"

critic:
	@sh -c "'$(CURDIR)/scripts/gocritic_check.sh'"

.PHONY: tests unittest golangci-lint critic