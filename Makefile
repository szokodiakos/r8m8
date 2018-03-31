build:
	go build

test:
	go test ./...

deps:
	dep ensure

infra:
	docker-compose up

default: build