
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: install
install:
	@echo "Install dependencies..."
	@go mod tidy
	@go mod vendor