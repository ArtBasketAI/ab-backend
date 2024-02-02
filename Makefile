GIT_REF := $(shell git describe --always)
VERSION ?= commit-$(GIT_REF)

SERVICE_NAME := $(shell grep "^module" go.mod | rev | cut -d "/" -f1 | rev)

# For pushing built docker images to public docker hub
DOCKER_HUB_IMAGE := bharathts07/$(SERVICE_NAME):$(VERSION)

.PHONY: build
build: build/server

build/server:
	CGO_ENABLED=0 go build -o bin/server \
	-ldflags "-X main.version=$(VERSION) -X main.serviceName=$(SERVICE_NAME)" \
	./cmd/server


.PHONY: run
run:
	@go run cmd/server/main.go

.PHONY: test
test: unit-test

.PHONY: unit-test
unit-test:
	@go test -count=1 -race -v $(shell go list ./...)

# GITHUB_TOKEN needed if the image needs to be pulled from a private repository
.PHONY: container
container:
	@docker build -t $(DOCKER_HUB_IMAGE) \
		--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
		--build-arg VERSION=$(VERSION) \
		--build-arg SERVICE_NAME=$(SERVICE_NAME) \
		--file Dockerfile \
		.

.PHONY: run-container
run-container:
	@echo "Forwading traffic from local port 80 to docker port 8080 (debugging)"
	@docker run -p 80:8080 $(DOCKER_HUB_IMAGE)

# This requires `docker login`
.PHONY: push-container
push-container:
	@docker push $(DOCKER_HUB_IMAGE)

.PHONY: dependencies
dependencies:
	@go mod tidy

.PHONY: lint
lint:
	golangci-lint run $(args) ./...
	go-consistent $(cons_args) ./...

.PHONY: lint-fix
lint-fix:
	@make lint args='--fix -v' cons_args='-v'
