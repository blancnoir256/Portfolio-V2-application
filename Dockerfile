# フロントエンドのビルドステージ
FROM node:20-alpine AS frontend-builder
WORKDIR /app

# パッケージをインストール
COPY ../frontend/package*.json ./
RUN npm ci

# ソースコードをコピーしてビルド
COPY ../frontend/ .
RUN npm run build


# バックエンドのビルドステージ
FROM golang:1.21-alpine AS backend-builder
WORKDIR /app

# 依存関係をインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server .


# 実行用の最終イメージ
FROM alpine:3.18
WORKDIR /app

# 必要なパッケージをインストール
RUN apk --no-cache add ca-certificates tzdata

# タイムゾーン設定
ENV TZ=Asia/Tokyo

# 必要なファイルをコピー
COPY --from=backend-builder /app/server ./
COPY --from=frontend-builder /app/dist ./public

# セキュリティのためnon-rootユーザーを作成
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser && \
    chown -R appuser:appgroup /app

# 環境変数を設定
ENV APP_DIR=/app
ENV SECRET_DIR=/app/secrets

# 必要なディレクトリを作成
RUN mkdir -p /app/secrets && \
    chown -R appuser:appgroup /app/secrets

# non-rootユーザーに切り替え
USER appuser

# ポート設定
EXPOSE 8080

# サーバー起動
CMD ["./server"]
