# Builder stage
FROM golang:1.24.2-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /workspace

# モジュールディレクトリを丸ごとコピー
COPY . .

# 依存関係を取得
RUN go mod download

# app モジュールをビルド
RUN go build -o /usr/local/bin/app-server ./app/cmd/server/main.go

# 実行ステージ
FROM golang:1.24.2-alpine AS runner

RUN apk add --no-cache git

# 作業ディレクトリを設定
WORKDIR /app

# ビルド成果物をコピー
COPY --from=builder /usr/local/bin/app-server /usr/local/bin/app-server

# マイグレーションファイルをコピー（コード内の相対パスに合わせて配置）
COPY app/infrastructure/db/migrations /app/infrastructure/db/migrations

# ポート公開
EXPOSE 8080

# エントリポイント
ENTRYPOINT ["app-server"]