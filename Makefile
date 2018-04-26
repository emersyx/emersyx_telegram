emtg.so: goget
	@go build -buildmode=plugin -o emtg.so emtg/*

.PHONY: goget
goget:
	@go get emersyx.net/emersyx/api
	@go get emersyx.net/emersyx/log
	@go get github.com/golang/lint/golint
	@go get github.com/BurntSushi/toml

.PHONY: test
test: emtg.so
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emtg/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
