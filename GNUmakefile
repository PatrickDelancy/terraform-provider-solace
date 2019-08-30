default: check test build

tools: ## Install the tools used to test and build
	@echo "==> Installing build tools"
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get github.com/goreleaser/goreleaser
	go get -u github.com/golang/dep/cmd/dep

build: ## Build for development purposes
	@echo "==> Running $@..."
	go build

test: ## Run the test suite with coverage
	@echo "==> Running $@..."
	@go test -cover -v -tags "$(BUILDTAGS)" \
		-race $(shell go list ./... | grep -v vendor)

testacc: ## Run the acceptance test suite with coverage
	@echo "==> Running $@..."
	@TF_ACC=1 go test -cover -v -tags "$(BUILDTAGS)" \
		-race $(shell go list ./... | grep -v vendor)

release: ## Trigger the release build script
	@echo "==> Running $@..."
	@goreleaser --rm-dist --config=goreleaser.yml

.PHONY: check
check: ## Run the gometalinter suite
	@echo "==> Running $@..."
	dep ensure
	golangci-lint run

HELP_FORMAT="    \033[36m%-25s\033[0m %s\n"
.PHONY: help
help: ## Display this usage information
	@echo "terraform-provider-solace make commands:"
	@grep -E '^[^ ]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
{printf $(HELP_FORMAT), $$1, $$2}'