
.PHONY: test
test: lint ## Run tests and create coverage report
	go test -short -coverprofile=coverage.txt -covermode=atomic ./... \
		&& go tool cover -func=coverage.txt

.PHONY: lint
lint: ## Run the code linter
	revive ./...

.PHONY: clean
clean: ## Clean up temp files and binaries
	rm -rf coverage*

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |sort \
		|awk 'BEGIN{FS=":.*?## "};{printf "\033[36m%-30s\033[0m %s\n",$$1,$$2}'
