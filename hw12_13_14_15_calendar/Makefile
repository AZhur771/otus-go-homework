BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

generate:
	protoc --go_out=api/stubs \
		  --go-grpc_out=api/stubs \
		  --grpc-gateway_out=api/stubs \
		  --grpc-gateway_opt=allow_delete_body=true \
		  --grpc-gateway_opt generate_unbound_methods=true \
  		  --proto_path=third_party/ \
  		  --proto_path=api/ \
  		  --openapiv2_out=third_party/openapiv2 \
  		  --openapiv2_opt allow_delete_body=true \
		  api/event_service.proto

protoc-version:
	which protoc && protoc --version

run: build
	$(BIN) -config ./configs/config.toml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

migrate:
	goose -dir=migrations postgres "user=postgres password=postgres dbname=calendar sslmode=disable" up

go-version:
	which go && go version

.PHONY: build run build-img run-img version test lint migrate
