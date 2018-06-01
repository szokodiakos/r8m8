build:
	go build

test:
	go test -timeout 30s ./...

deps:
	dep ensure

infra:
	docker-compose up

default: build
