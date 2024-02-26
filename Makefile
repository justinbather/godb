DOCKER_IMG?=justinbather/godb
BIN_DIR?=./bin
TAG?=$(shell git rev-parse --short HEAD)

.PHONY: test
test:
	go test -v -race ./...

.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: godb
godb:
	mkdir -p $(BIN_DIR)
	go mod tidy
	CGO_ENABLED=0 go build -o $(BIN_DIR)/example ./cmd/example

.PHONY: docker-build
docker-build:
	docker build -t=$(DOCKER_IMG):$(TAG) .

.PHONY: docker-run
docker-run:
	docker run --rm -ti $(DOCKER_IMG):$(TAG)

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_IMG):$(TAG)

.PHONY: server
server:
	mkdir -p $(BIN_DIR)
	go mod tidy
	CGO_ENABLED=0 go build -o $(BIN_DIR)/server ./cmd/server
