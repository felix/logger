
GOPATH?=	$(HOME)/go

.PHONY: test
test: lint ## Run tests and create coverage report
	go test -short -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

.PHONY: lint
lint: $(GOPATH)/bin/golint ## Run the code linter
	@for file in $$(find . -name 'vendor' -prune -o -type f -name '*.go'); do \
		golint $$file; done

$(GOPATH)/bin/golint:
	go get -u golang.org/x/lint/golint

.PHONY: clean
clean: ## Clean up temp files and binaries
	@rm -rf coverage*

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |sort \
		|awk 'BEGIN{FS=":.*?## "};{printf "\033[36m%-30s\033[0m %s\n",$$1,$$2}'
