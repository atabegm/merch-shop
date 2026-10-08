include .env

.PHONY: build test migrate-create migrate-up migrate-down docker-up docker-build docker-down mocks
.DEFAULT_GOAL: test

test: 
	go test -v -race -timeout 30s ./...

build:
	go build -v -o ./bin/app ./cmd/app
	go build -v -0 ./bin/consumer ./cmd/consumer

migrate-create: 
	migrate create -ext sql -dir migrations second migration

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1
	
docker-up:
	docker compose up -d 

docker-build:
	docker compose up -d --build

docker-down:
	docker compose down

mocks-service:
	mockgen -source=internal/service/contracts.go -destination=internal/service/mocks/mocks.go -package=mocks

mocks-api:
	mockgen -source=internal/api/contracts.go -destination=internal/api/mocks/mocks.go -package=mocks
