# Этап 1: Сборка фронтенда
FROM node:18-alpine AS frontend-builder

WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend .

RUN npm run build

# Проверим что билд есть
RUN ls -la dist

# Этап 2: Сборка Go-приложения
FROM golang:1.21-alpine AS go-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Копируем фронтенд билд
COPY --from=frontend-builder /frontend/dist ./public

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /feedbox ./cmd/http/main.go

# Финальный образ
FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY --from=go-builder /feedbox .
COPY --from=go-builder /app/local.yaml .
COPY --from=go-builder /app/internal ./internal/config
COPY --from=go-builder /app/public ./public

RUN chmod +x /app/feedbox

ENV CONFIG_PATH=/app/local.yaml

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD wget -qO- http://localhost:8080/api/health || exit 1

CMD ["./feedbox"]