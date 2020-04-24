
.PHONY: test
test: lint
	go test -short -coverprofile=coverage.txt -covermode=atomic ./... \
		&& go tool cover -func=coverage.txt

.PHONY: lint
lint:
	go vet ./...

.PHONY: clean
clean:
	rm -rf coverage*
