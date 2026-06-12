# LeetCode (Go) — local practice shortcuts.
# Requires: Go (see go.mod). Optional: golangci-lint, gofumpt.

SHELL       := /usr/bin/env bash
.SHELLFLAGS := -eu -o pipefail -c

# --- go ---
GO          ?= go
TESTFLAGS   ?= -count=1
COVER_OUT   := coverage.out

# Problem filter for test/bench/cover. Matches the import path
# case-insensitively, so P=0001 or P=two-sum both work.
P           ?=

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z0-9_.-]+:.*?##' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

# --- scaffold ---
.PHONY: new
new: ## Scaffold a problem: make new ID=1 SLUG=two-sum [KIND=plain|list|tree] [TITLE="Two Sum"] [DIFF=easy|medium|hard] [URL=...]
	@ID="$(ID)" SLUG="$(SLUG)" KIND="$(KIND)" TITLE="$(TITLE)" DIFF="$(DIFF)" URL="$(URL)" bash scripts/new.sh

# --- test ---
# Resolve the package set once: all problems, or those matching P.
define _pkgs
$(shell if [ -n "$(P)" ]; then $(GO) list ./problems/... 2>/dev/null | grep -i -- "$(P)"; else echo ./...; fi)
endef

.PHONY: test
test: ## Run tests (filter one problem with P=<id-or-slug>, e.g. make test P=0001)
	@pkgs="$(call _pkgs)"; \
	if [ -z "$$pkgs" ]; then echo "no problems match P='$(P)'"; exit 1; fi; \
	echo "go test $$pkgs"; \
	$(GO) test $(TESTFLAGS) $$pkgs

.PHONY: test-race
test-race: ## Run tests with the race detector
	@$(GO) test $(TESTFLAGS) -race $(call _pkgs)

.PHONY: bench
bench: ## Run benchmarks (filter with P=<id-or-slug>)
	@$(GO) test $(TESTFLAGS) -run=^$$ -bench=. -benchmem $(call _pkgs)

.PHONY: cover
cover: ## Run tests with coverage and open the HTML report
	@$(GO) test $(TESTFLAGS) -coverprofile=$(COVER_OUT) $(call _pkgs)
	@$(GO) tool cover -html=$(COVER_OUT)

# --- quality ---
.PHONY: fmt
fmt: ## Format code (gofumpt if installed, else gofmt)
	@if command -v gofumpt >/dev/null 2>&1; then gofumpt -w .; else $(GO) fmt ./...; fi

.PHONY: vet
vet: ## Run go vet
	@$(GO) vet ./...

.PHONY: lint
lint: ## Run golangci-lint if installed
	@if command -v golangci-lint >/dev/null 2>&1; then golangci-lint run; else echo "golangci-lint not installed — skipping"; fi

.PHONY: tidy
tidy: ## Tidy go.mod
	@$(GO) mod tidy

.PHONY: verify
verify: fmt vet test ## fmt, vet, then run all tests

# --- info ---
.PHONY: list
list: ## List all problems
	@find problems -mindepth 1 -maxdepth 1 -type d | sort | sed 's:problems/::'

.PHONY: count
count: ## Count solved problems
	@n=$$(find problems -mindepth 1 -maxdepth 1 -type d | wc -l | tr -d ' '); echo "$$n problems"

.PHONY: clean
clean: ## Remove coverage artifacts
	@rm -f $(COVER_OUT)
