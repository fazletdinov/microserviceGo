.PHONY: proto
proto:
	protoc -I ./proto \
  --go_out=./protogen/golang --go_opt=paths=source_relative \
  --go-grpc_out=./protogen/golang --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./protogen/golang --grpc-gateway_opt=paths=source_relative \
  --openapiv2_out=./protogen/openapiv2 \
  --openapiv2_opt=logtostderr=true \
  --openapiv2_opt=generate_unbound_methods=true \
  --openapiv2_opt=allow_merge=true \
  --openapiv2_opt=merge_file_name=allservices \
  ./proto/auth/auth.proto ./proto/auth/auth_service.proto ./proto/likes/likes.proto \
  ./proto/likes/likes_service.proto ./proto/posts/posts.proto ./proto/posts/posts_service.proto \

.PHONY: cert-jwt
cert-jwt:
	openssl genrsa -out cert/id_rsa 4096
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub

.PHONY: start
start:
	docker compose -f docker-compose.gateway.yaml up --build

.PHONY: cert-jaeger
cert-jaeger:
	openssl genrsa -out cert/jaeger/server.key 2048
	openssl req -new -key cert/jaeger/server.key -out cert/jaeger/cert.crt