FROM golang:1.25.1-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go env -w GOPROXY=https://proxy.golang.org,direct
COPY . .
RUN go build -o /app/life_manager ./cmd/main.go


FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/life_manager /app/life_manager
EXPOSE 8080
CMD ["/app/life_manager"]