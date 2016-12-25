.PHONY: all prepare build vet lint test

all: prepare vet lint test build

prepare:
	go get -v github.com/golang/lint/golint
	go get -v github.com/Masterminds/glide
	glide install

build:
	GOARCH=amd64 GOOS=linux go install

vet:
	GOARCH=amd64 GOOS=linux go vet $$(go list ./... | grep -v /vendor/)

lint:
	for pkg in $$(go list ./... | grep -v /vendor/); do golint $$pkg; done

test:
	GOARCH=amd64 GOOS=linux go test $$(go list ./... | grep -v /vendor/)
