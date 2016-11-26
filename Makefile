all: test build

build:
	GOARCH=amd64 GOOS=linux go install

test:
	GOARCH=amd64 GOOS=linux go test $$(go list ./... | grep -v /vendor/ | grep -v /Godeps/)
