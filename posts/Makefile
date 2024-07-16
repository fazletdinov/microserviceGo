include .env

.PHONY: migrate
migrate:
	goose -dir ./migrations postgres "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_NAME}?sslmode=disable" up

.PHONY: proto
proto:
	protoc -I ./proto \
  --go_out=./protogen --go_opt=paths=source_relative \
  --go-grpc_out=./protogen --go-grpc_opt=paths=source_relative \
  ./proto/posts/posts.proto ./proto/posts/posts_service.proto