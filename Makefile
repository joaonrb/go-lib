PROJECT_NAME := "go-lib"
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)


.PHONY: \
	get-dependencies \
	test \
	vtest \
	format-code \
    lint \
    install-tools \
	help

get-dependencies: ## Get binary source dependencies
	@go mod download

test: ## Run unit-test
	@go test -coverpkg=./... -race -short -coverprofile cov.out ./...

vtest: ## Run verbose unit-test
	@go test -coverpkg=./... -race -short -coverprofile cov.out -v ./...

coverage: ## Show coverage in html
	@go tool cover -html=cov.out

format-code:  ## format the code
	@goimports -w .
	@golines -w .
	@go fmt ./...

lint:  ## Lint the code
	@golangci-lint run

install-tools: ## Install all the dependencies for the development
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
