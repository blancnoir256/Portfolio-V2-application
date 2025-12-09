# ---------- Frontend build ----------
FROM node:20-slim AS frontend-builder
WORKDIR /app/web

# 依存解決
COPY ./web/package*.json ./
ENV CI=true
RUN npm ci

# ソースコピー＆ビルド
COPY ./web/ .
RUN npm run build

# ---------- Backend build ----------
FROM golang:1.25-alpine AS backend-router-builder
WORKDIR /app/backend-router

# 依存取得
COPY ./go.mod ./go.sum ./
RUN go mod download

# ソースコピー
COPY ./main.go ./
COPY ./api/ ./api/
COPY ./router/ ./router/

# 静的ビルド (必要に応じて -s -w)
ENV CGO_ENABLED=0
RUN go build -trimpath -ldflags="-s -w" -o server .

# ---------- Runtime image ----------
FROM alpine:3.18
WORKDIR /app

ENV TZ=Asia/Tokyo \
    APP_PORT=8080 \
    DIST_DIR=/app/public

# タイムゾーンと証明書
RUN apk add --no-cache tzdata ca-certificates && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && \
    addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

# コピーして持ってくる
COPY --from=backend-router-builder --chown=appuser:appuser /app/backend-router/server ./server
COPY --from=frontend-builder    --chown=appuser:appuser /app/web/dist ./public
COPY --from=frontend-builder    --chown=appuser:appuser /app/web/public ./public

USER appuser

EXPOSE ${APP_PORT}

CMD ["./server"]
