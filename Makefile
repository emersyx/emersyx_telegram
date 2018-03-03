emtg.so: goget
	go build -buildmode=plugin -o emtg.so emtg/*

.PHONY: goget
goget:
	go get emersyx.net/emersyx_apis/emcomapi
	go get emersyx.net/emersyx_apis/emtgapi

.PHONY: test
test: emtg.so
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emtg/*.go tgbotapi/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
