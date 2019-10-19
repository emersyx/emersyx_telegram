telegram.so:
	@go build -buildmode=plugin -o telegram.so internal/telegram/*

.PHONY: test
test: telegram.so
	@echo "Running the tests with gofmt..."
	@test -z $(shell gofmt -s -l internal/telegra/*.go)
	@echo "Running the tests with go vet..."
	@go vet ./...
	@echo "Running the tests with golint..."
	@golint -set_exit_status $(shell go list ./...)
