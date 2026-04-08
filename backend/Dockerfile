FROM golang:1.25.0-alpine AS builder

WORKDIR /workspace

COPY . .

RUN go mod tidy

RUN go build -o InfinityShortcut-Backend ./cmd/main.go

RUN chmod +x InfinityShortcut-Backend

FROM alpine:latest

WORKDIR /app

COPY --from=builder /workspace/InfinityShortcut-Backend .

EXPOSE 8888

ENTRYPOINT [ "/app/InfinityShortcut-Backend" ]