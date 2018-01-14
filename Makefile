.PHONY: emtg-dep emtg-goget test

emtg-dep:
	dep ensure
	go build -buildmode=plugin -o emtg.so emtg/*

emtg-goget:
	go get -t -buildmode=plugin ./emtg

test:
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emtg/*.go tgbotapi/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
