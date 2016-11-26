.PHONY: prepare build vet test

all: vet test build

prepare:
	go get github.com/Masterminds/glide
	glide install

build:
	GOARCH=amd64 GOOS=linux go install

vet:
	GOARCH=amd64 GOOS=linux go vet $$(go list ./... | grep -v /vendor/)

test:
	GOARCH=amd64 GOOS=linux go test $$(go list ./... | grep -v /vendor/)
