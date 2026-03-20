.PHONY: build test test-unit test-all cover cover-html bench lint fmt clean check-gh branch pr help

BINARY := multi-git-mirror
GOFLAGS := -v

## Build

build: ## Build the binary
	go build $(GOFLAGS) -o $(BINARY) ./cmd/main.go

## Test

test: test-unit ## Run unit tests (alias)

test-unit: ## Run unit tests with coverage
	go test ./internal/... ./cmd/... -v -race -cover

test-all: test-unit ## Run all tests

## Coverage

cover: ## Generate coverage report
	go test ./internal/... ./cmd/... -coverprofile=coverage.out
	go tool cover -func=coverage.out

cover-html: cover ## Open coverage report in browser
	go tool cover -html=coverage.out

## Benchmark

bench: ## Run benchmarks
	go test -bench=. -benchmem ./internal/...

## Quality

lint: ## Run go vet
	go vet ./...

fmt: ## Format code
	gofmt -s -w .

## Cleanup

clean: ## Remove build artifacts and coverage files
	rm -f $(BINARY) coverage.out

## Workflow

check-gh: ## Check if gh CLI is installed and authenticated
	@command -v gh >/dev/null 2>&1 || { echo "\033[31m✗ gh CLI not installed. Run: brew install gh\033[0m"; exit 1; }
	@gh auth status >/dev/null 2>&1 || { echo "\033[31m✗ gh CLI not authenticated. Run: gh auth login\033[0m"; exit 1; }
	@echo "\033[32m✓ gh CLI ready\033[0m"

branch: ## Create feature branch (usage: make branch name=watch-mode)
	@if [ -z "$(name)" ]; then echo "Usage: make branch name=<feature-name>"; exit 1; fi
	git checkout main
	git pull origin main
	git checkout -b feat/$(name)
	@echo "\033[32m✓ Branch feat/$(name) created\033[0m"

pr: check-gh ## Run tests, push, and create PR (usage: make pr title="Add feature")
	@if [ -z "$(title)" ]; then echo "Usage: make pr title=\"PR title\""; exit 1; fi
	go test ./... -race -cover
	go vet ./...
	git push -u origin $$(git branch --show-current)
	gh pr create --title "$(title)" --body "## Summary"$$'\n\n'"Branch: $$(git branch --show-current)"$$'\n\n'"## Test plan"$$'\n\n'"- [ ] Unit tests pass"$$'\n'"- [ ] Coverage maintained"
	@echo "\033[32m✓ PR created\033[0m"

## Help

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
