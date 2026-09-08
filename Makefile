.PHONY: fmt test vet build run compose-up compose-down verify

fmt:
	gofmt -w cmd internal

test:
	go test -race -cover ./...

vet:
	go vet ./...

build:
	CGO_ENABLED=0 go build -trimpath -o bin/streaming-api ./cmd/api

run:
	go run ./cmd/api

compose-up:
	docker compose up --build -d

compose-down:
	docker compose down --remove-orphans

verify: test vet

