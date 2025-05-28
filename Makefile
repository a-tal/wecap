lint:
	golangci-lint run ./...

test: lint
	go test -cover ./...

build:
	CGO_ENABLED=0 go build .

docker:
	docker build -t wecap:dev .

.PHONY: lint test build docker
