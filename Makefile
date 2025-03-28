DATE := $(shell date +%Y)

##@ Helpers
.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1;34m${DOCKER_NAMESPACE}\033[0m\tCopyright (c) ${DATE} XXXX Development\n \n\033[1;32mUsage:\033[0m\n  make \033[1;34m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[1;34m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Basic
.PHONY: test
test: ## is used to run the test suite of the application
	@go test ./...

.PHONY: test-verbose
test-verbose: ## is used to run the test suite of the application in verbose mode
	@go test ./... -v

.PHONY: coverage
coverage: test ## is used to generate the coverage report of the application
	@go clean -testcache; go test ./... -coverprofile=coverage.out; go tool cover -func=coverage.out

.PHONY: format
format: ## is used to format the code of the application
	@gofmt -d .
