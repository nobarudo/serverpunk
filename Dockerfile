# ==========================================
# Stage 1: フロントエンドのビルド (Node.js)
# ==========================================
FROM --platform=$BUILDPLATFORM node:24-bookworm-slim AS frontend-builder
WORKDIR /app/front

RUN corepack enable && corepack prepare pnpm@latest --activate

COPY front/package.json front/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

# フロントエンドのソースコードをコピーしてビルド
COPY front/ ./
RUN pnpm build

# ==========================================
# Stage 2: バックエンドのビルド (Go + CGO)
# ==========================================
FROM --platform=linux/amd64 golang:1.26-bookworm AS backend-builder
WORKDIR /app/back

# ★ 先にgo.modだけコピーして依存関係を解決
COPY back/go.mod back/go.sum* ./
RUN go mod download

# Goのソースコード全体をコピー
COPY back/ ./

COPY --from=frontend-builder /app/front/build ./internal/platform/server/build

# ビルド実行
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -p 1 -o serverpunk cmd/server/main.go

# ==========================================
# Stage 3: 実行環境 (Runtime)
# ==========================================
# 本番環境に合わせて極小のRocky Linuxベースを使用
FROM --platform=linux/amd64 rockylinux:9-minimal
WORKDIR /app

# Stage2で錬成された「Svelte内包の単一バイナリ」だけを抽出
COPY --from=backend-builder /app/back/serverpunk .

# 実行権限の付与
RUN chmod +x ./serverpunk

# Ginのデフォルトポートを開放
EXPOSE 8080

# サーバーちゃんの起動！
CMD ["./serverpunk"]
