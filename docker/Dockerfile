FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /feedbox ./cmd/http/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /feedbox .
COPY --from=builder /app/local.yaml .
COPY --from=builder /app/config ./internal/config

ENV CONFIG_PATH=/app/local.yaml  

CMD ["./feedbox"]
