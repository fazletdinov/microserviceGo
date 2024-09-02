FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata
# RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o main .


FROM golang:1.22-alpine

# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /build/main /app/main
COPY --from=builder /build/config /app/config
COPY --from=builder /build/.env /app/.env
COPY --from=builder /build/cert /app/cert

CMD ["./main"]