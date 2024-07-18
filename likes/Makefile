include .env

.DEFAULT_GOAL := build

.PHONY: proto
proto:
	protoc -I ./proto \
  --go_out=./protogen --go_opt=paths=source_relative \
  --go-grpc_out=./protogen --go-grpc_opt=paths=source_relative \
  ./proto/likes/likes.proto ./proto/likes/likes_service.proto


.PHONY: migrate
migrate:
	goose -dir ./migrations postgres "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_NAME}?sslmode=disable" up

.PHONY: build
build: migrate
	go build -o main .

.PHONY: start
start:
	docker compose -f docker-compose.likes.yaml up --build